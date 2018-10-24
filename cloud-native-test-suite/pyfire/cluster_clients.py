import json
import subprocess
import requests

NUM_MAX_PROJECTS=10000000

class Client:
    def k8sCommand():
        pass

''' 
K8S Client 
'''
class K8sClient:
    def __init__(self):
        pass

    def get_namespaces(self):
        # K8s command to get a table of namespaces without the header line
        ns, err = subprocess.Popen(["oc","get", "ns","--no-headers"], stdout=subprocess.PIPE).communicate()
        # Split output by row
        ns_list = ns.split("\n")[:-1] # remove empty string as last entry
        # Get first entry in each row
        ns_list = [n.split()[0] for n in ns_list]
        return ns_list
        

    def get_images(self, namespace=""):
        # Select all namespaces if one wasn't provided
        namespace = '--all-namespaces' if namespace == '' else '-n '+namespace 
        # Path to an image within a Pod's spec
        image_path = '{.items[*].spec.containers[*].image}'
        # K8s Terminal Command to get all pods and extract the images from the pod's json
        k8s_command = ''.join(['oc get pods ',namespace,' -o jsonpath="',image_path,'"'])
        images, err = subprocess.Popen(k8s_command, shell=True, stdout=subprocess.PIPE).communicate()
        # Remove initial "
        return images[1:].split()

    def get_images_per_pod(self, namespace=""):
        namespace = '--all-namespaces' if namespace == '' else '-n '+namespace 
        k8s_command = ''.join(["oc get pods ",namespace," -o json"])
        pods = json.loads(subprocess.Popen(k8s_command, shell=True, stdout=subprocess.PIPE).communicate()[0])
        pod_dump = []
        for pod in pods['items']:
            ns = pod['metadata']['namespace']
            try:
                for container in pod['status']['containerStatuses']:
                    name = container['name']
                    image = container['image']
                    imageID = container['imageID']
                    pod_dump.append([ns, name, image, imageID])
            except:
                print("======== NO CONTAINER STATUSES ========")
                print(json.dumps(pod, indent=2))
        return pod_dump

    def get_annotations_per_pod(self, namespace=""):
        namespace = '--all-namespaces' if namespace == '' else '-n '+namespace 
        k8s_command = ''.join(["oc get pods ",namespace," -o json"])
        pods = json.loads(subprocess.Popen(k8s_command, shell=True, stdout=subprocess.PIPE).communicate()[0])
        pod_dump = []
        for pod in pods['items']:
            ns = pod['metadata']['namespace']
            name = pod['metadata']['name']
            annotation = pod['metadata']['annotations']
            pod_dump.append([ns, name, annotation])
        return pod_dump

    def get_labels_per_pod(self, namespace=""):
        namespace = '--all-namespaces' if namespace == '' else '-n '+namespace 
        k8s_command = ''.join(["oc get pods ",namespace," -o json"])
        pods = json.loads(subprocess.Popen(k8s_command, shell=True, stdout=subprocess.PIPE).communicate()[0])
        pod_dump = []
        for pod in pods['items']:
            ns = pod['metadata']['namespace']
            name = pod['metadata']['name']
            annotation = pod['metadata']['labels']
            pod_dump.append([ns, name, annotation])
        return pod_dump

    def get_all_projects(self): 
        # K8s command to get all projects
        projects, err = subprocess.Popen(["oc","get", "projects", "--no-headers"], stdout=subprocess.PIPE).communicate()
        projects_list = ns.split("\n")[:-1] # remove empty string as last entry
        ns_list = [n.split()[0] for n in ns_list]
        return projects.split()

'''
Hub Client
'''
class HubClient:
    def __init__(self, host_name):
        self.host_name = host_name
        self.secure_login_cookie = self.get_secure_login_cookie()

    def get_secure_login_cookie(self):
        security_headers = {'Content-Type':'application/x-www-form-urlencoded'}
        security_data = {'j_username':'sysadmin','j_password':'blackduck'}
        # verify=False does not verify SSL connection - insecure
        r = requests.post("https://"+self.host_name+":443/j_spring_security_check", verify=False, data=security_data, headers=security_headers)
        return r.cookies 
        
    def get_projects_dump(self): 
        r = requests.get("https://"+self.host_name+":443/api/projects?limit="+str(NUM_MAX_PROJECTS),verify=False, cookies=self.secure_login_cookie)
        return r.json()['items']

    def get_projects_names(self):
        return [x['name'] for x in self.get_projects_dump()]

    def get_code_locations_dump(self):
        r = requests.get("https://"+self.host_name+":443/api/codelocations?limit="+str(NUM_MAX_PROJECTS),verify=False, cookies=self.secure_login_cookie)
        return r.json()

    def get_code_locations_names(self):
        return [x['name'] for x in self.get_code_locations_dump()['items']]

'''
OpsSight Client
'''

class OpsSightClient:
    def __init__(self, host_name):
        self.host_name = host_name

    def create(self):
        pass # spin up with Matt's yaml
        # get the url by exposing 
    
    def get_dump(self):
        r = requests.get("http://"+self.host_name+"/model")
        return json.loads(r.text)

    def get_shas_names(self):
        return self.get_dump()['CoreModel']['Images'].keys()