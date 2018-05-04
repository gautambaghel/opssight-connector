/*
Copyright (C) 2018 Synopsys, Inc.

Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements. See the NOTICE file
distributed with this work for additional information
regarding copyright ownership. The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied. See the License for the
specific language governing permissions and limitations
under the License.
*/

package annotations

import (
	"fmt"
	"strings"

	"github.com/blackducksoftware/perceivers/pkg/annotations"
)

// ImageAnnotationPrefix is the prefix used for BlackDuckAnnotations in image annotations
var ImageAnnotationPrefix = "quality.image.openshift.io"

// CreateImageLabels returns a map of labels from a ImageAnnotationData object
func CreateImageLabels(obj interface{}, name string, count int) map[string]string {
	imageData := obj.(*annotations.ImageAnnotationData)
	imagePostfix := ""
	labels := make(map[string]string)

	if len(name) > 0 {
		imagePostfix = fmt.Sprintf("%d", count)
		labels[fmt.Sprintf("com.blackducksoftware.image%d", count)] = ShortenLabelContent(name)
	}
	labels[fmt.Sprintf("com.blackducksoftware.image%s.policy-violations", imagePostfix)] = fmt.Sprintf("%d", imageData.GetPolicyViolationCount())
	labels[fmt.Sprintf("com.blackducksoftware.image%s.has-policy-violations", imagePostfix)] = fmt.Sprintf("%t", imageData.HasPolicyViolations())
	labels[fmt.Sprintf("com.blackducksoftware.image%s.vulnerabilities", imagePostfix)] = fmt.Sprintf("%d", imageData.GetVulnerabilityCount())
	labels[fmt.Sprintf("com.blackducksoftware.image%s.has-vulnerabilities", imagePostfix)] = fmt.Sprintf("%t", imageData.HasVulnerabilities())
	labels[fmt.Sprintf("com.blackducksoftware.image%s.overall-status", imagePostfix)] = imageData.GetOverallStatus()

	return labels
}

// CreateImageAnnotations returns a map of annotations from a ImageAnnotationData object
func CreateImageAnnotations(obj interface{}, name string, count int) map[string]string {
	imageData := obj.(*annotations.ImageAnnotationData)
	imagePrefix := ""
	newAnnotations := make(map[string]string)

	if len(name) > 0 {
		imagePrefix = fmt.Sprintf("image%d.", count)
		imageName := strings.Replace(name, "/", ".", -1)
		newAnnotations[fmt.Sprintf("%sblackducksoftware.com", imagePrefix)] = imageName
		newAnnotations[fmt.Sprintf("%s%s", imagePrefix, ImageAnnotationPrefix)] = imageName
	}
	newAnnotations[fmt.Sprintf("%sblackducksoftware.com/hub-scanner-version", imagePrefix)] = imageData.GetScanClientVersion()
	newAnnotations[fmt.Sprintf("%sblackducksoftware.com/attestation-server-version", imagePrefix)] = imageData.GetServerVersion()
	newAnnotations[fmt.Sprintf("%sblackducksoftware.com/project-endpoint", imagePrefix)] = imageData.GetComponentsURL()

	vulnAnnotations := CreateBlackDuckVulnerabilityAnnotation(imageData.HasVulnerabilities() == true, imageData.GetComponentsURL(), imageData.GetVulnerabilityCount(), imageData.GetScanClientVersion())
	policyAnnotations := CreateBlackDuckPolicyAnnotation(imageData.HasPolicyViolations() == true, imageData.GetComponentsURL(), imageData.GetPolicyViolationCount(), imageData.GetScanClientVersion())

	newAnnotations[fmt.Sprintf("%s%s/vulnerability.blackduck", imagePrefix, ImageAnnotationPrefix)] = vulnAnnotations.AsString()
	newAnnotations[fmt.Sprintf("%s%s/policy.blackduck", imagePrefix, ImageAnnotationPrefix)] = policyAnnotations.AsString()

	return newAnnotations
}

// ShortenLabelContent will ensure the data is less than the 63 character limit and doesn't contain
// any characters that are not allowed
func ShortenLabelContent(data string) string {
	newData := RemoveRegistryInfo(data)

	// Label values can not be longer than 63 characters
	if len(newData) > 63 {
		newData = newData[0:63]
	}

	return newData
}

// RemoveRegistryInfo will take a string and return a string that removes any registry name information
// and replaces all / with .
func RemoveRegistryInfo(d string) string {
	s := strings.Split(d, "/")

	// If the data includes a . or : before the first / then that string is most likely
	// a registry name.  Remove it because it could make the data too long and
	// truncate useful information
	if strings.Contains(s[0], ".") || strings.Contains(s[0], ":") {
		s = s[1:]
	}
	return strings.Join(s, ".")
}
