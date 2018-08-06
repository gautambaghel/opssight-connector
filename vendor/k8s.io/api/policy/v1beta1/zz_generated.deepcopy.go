// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	core_v1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedFlexVolume) DeepCopyInto(out *AllowedFlexVolume) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedFlexVolume.
func (in *AllowedFlexVolume) DeepCopy() *AllowedFlexVolume {
	if in == nil {
		return nil
	}
	out := new(AllowedFlexVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedHostPath) DeepCopyInto(out *AllowedHostPath) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedHostPath.
func (in *AllowedHostPath) DeepCopy() *AllowedHostPath {
	if in == nil {
		return nil
	}
	out := new(AllowedHostPath)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Eviction) DeepCopyInto(out *Eviction) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.DeleteOptions != nil {
		in, out := &in.DeleteOptions, &out.DeleteOptions
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.DeleteOptions)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Eviction.
func (in *Eviction) DeepCopy() *Eviction {
	if in == nil {
		return nil
	}
	out := new(Eviction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Eviction) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FSGroupStrategyOptions) DeepCopyInto(out *FSGroupStrategyOptions) {
	*out = *in
	if in.Ranges != nil {
		in, out := &in.Ranges, &out.Ranges
		*out = make([]IDRange, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FSGroupStrategyOptions.
func (in *FSGroupStrategyOptions) DeepCopy() *FSGroupStrategyOptions {
	if in == nil {
		return nil
	}
	out := new(FSGroupStrategyOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostPortRange) DeepCopyInto(out *HostPortRange) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostPortRange.
func (in *HostPortRange) DeepCopy() *HostPortRange {
	if in == nil {
		return nil
	}
	out := new(HostPortRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IDRange) DeepCopyInto(out *IDRange) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IDRange.
func (in *IDRange) DeepCopy() *IDRange {
	if in == nil {
		return nil
	}
	out := new(IDRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDisruptionBudget) DeepCopyInto(out *PodDisruptionBudget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodDisruptionBudget.
func (in *PodDisruptionBudget) DeepCopy() *PodDisruptionBudget {
	if in == nil {
		return nil
	}
	out := new(PodDisruptionBudget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodDisruptionBudget) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDisruptionBudgetList) DeepCopyInto(out *PodDisruptionBudgetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PodDisruptionBudget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodDisruptionBudgetList.
func (in *PodDisruptionBudgetList) DeepCopy() *PodDisruptionBudgetList {
	if in == nil {
		return nil
	}
	out := new(PodDisruptionBudgetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodDisruptionBudgetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDisruptionBudgetSpec) DeepCopyInto(out *PodDisruptionBudgetSpec) {
	*out = *in
	if in.MinAvailable != nil {
		in, out := &in.MinAvailable, &out.MinAvailable
		if *in == nil {
			*out = nil
		} else {
			*out = new(intstr.IntOrString)
			**out = **in
		}
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.LabelSelector)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		if *in == nil {
			*out = nil
		} else {
			*out = new(intstr.IntOrString)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodDisruptionBudgetSpec.
func (in *PodDisruptionBudgetSpec) DeepCopy() *PodDisruptionBudgetSpec {
	if in == nil {
		return nil
	}
	out := new(PodDisruptionBudgetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDisruptionBudgetStatus) DeepCopyInto(out *PodDisruptionBudgetStatus) {
	*out = *in
	if in.DisruptedPods != nil {
		in, out := &in.DisruptedPods, &out.DisruptedPods
		*out = make(map[string]v1.Time, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodDisruptionBudgetStatus.
func (in *PodDisruptionBudgetStatus) DeepCopy() *PodDisruptionBudgetStatus {
	if in == nil {
		return nil
	}
	out := new(PodDisruptionBudgetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSecurityPolicy) DeepCopyInto(out *PodSecurityPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSecurityPolicy.
func (in *PodSecurityPolicy) DeepCopy() *PodSecurityPolicy {
	if in == nil {
		return nil
	}
	out := new(PodSecurityPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodSecurityPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSecurityPolicyList) DeepCopyInto(out *PodSecurityPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PodSecurityPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSecurityPolicyList.
func (in *PodSecurityPolicyList) DeepCopy() *PodSecurityPolicyList {
	if in == nil {
		return nil
	}
	out := new(PodSecurityPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodSecurityPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSecurityPolicySpec) DeepCopyInto(out *PodSecurityPolicySpec) {
	*out = *in
	if in.DefaultAddCapabilities != nil {
		in, out := &in.DefaultAddCapabilities, &out.DefaultAddCapabilities
		*out = make([]core_v1.Capability, len(*in))
		copy(*out, *in)
	}
	if in.RequiredDropCapabilities != nil {
		in, out := &in.RequiredDropCapabilities, &out.RequiredDropCapabilities
		*out = make([]core_v1.Capability, len(*in))
		copy(*out, *in)
	}
	if in.AllowedCapabilities != nil {
		in, out := &in.AllowedCapabilities, &out.AllowedCapabilities
		*out = make([]core_v1.Capability, len(*in))
		copy(*out, *in)
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]FSType, len(*in))
		copy(*out, *in)
	}
	if in.HostPorts != nil {
		in, out := &in.HostPorts, &out.HostPorts
		*out = make([]HostPortRange, len(*in))
		copy(*out, *in)
	}
	in.SELinux.DeepCopyInto(&out.SELinux)
	in.RunAsUser.DeepCopyInto(&out.RunAsUser)
	in.SupplementalGroups.DeepCopyInto(&out.SupplementalGroups)
	in.FSGroup.DeepCopyInto(&out.FSGroup)
	if in.DefaultAllowPrivilegeEscalation != nil {
		in, out := &in.DefaultAllowPrivilegeEscalation, &out.DefaultAllowPrivilegeEscalation
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.AllowPrivilegeEscalation != nil {
		in, out := &in.AllowPrivilegeEscalation, &out.AllowPrivilegeEscalation
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.AllowedHostPaths != nil {
		in, out := &in.AllowedHostPaths, &out.AllowedHostPaths
		*out = make([]AllowedHostPath, len(*in))
		copy(*out, *in)
	}
	if in.AllowedFlexVolumes != nil {
		in, out := &in.AllowedFlexVolumes, &out.AllowedFlexVolumes
		*out = make([]AllowedFlexVolume, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSecurityPolicySpec.
func (in *PodSecurityPolicySpec) DeepCopy() *PodSecurityPolicySpec {
	if in == nil {
		return nil
	}
	out := new(PodSecurityPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunAsUserStrategyOptions) DeepCopyInto(out *RunAsUserStrategyOptions) {
	*out = *in
	if in.Ranges != nil {
		in, out := &in.Ranges, &out.Ranges
		*out = make([]IDRange, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunAsUserStrategyOptions.
func (in *RunAsUserStrategyOptions) DeepCopy() *RunAsUserStrategyOptions {
	if in == nil {
		return nil
	}
	out := new(RunAsUserStrategyOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SELinuxStrategyOptions) DeepCopyInto(out *SELinuxStrategyOptions) {
	*out = *in
	if in.SELinuxOptions != nil {
		in, out := &in.SELinuxOptions, &out.SELinuxOptions
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.SELinuxOptions)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SELinuxStrategyOptions.
func (in *SELinuxStrategyOptions) DeepCopy() *SELinuxStrategyOptions {
	if in == nil {
		return nil
	}
	out := new(SELinuxStrategyOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SupplementalGroupsStrategyOptions) DeepCopyInto(out *SupplementalGroupsStrategyOptions) {
	*out = *in
	if in.Ranges != nil {
		in, out := &in.Ranges, &out.Ranges
		*out = make([]IDRange, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SupplementalGroupsStrategyOptions.
func (in *SupplementalGroupsStrategyOptions) DeepCopy() *SupplementalGroupsStrategyOptions {
	if in == nil {
		return nil
	}
	out := new(SupplementalGroupsStrategyOptions)
	in.DeepCopyInto(out)
	return out
}
