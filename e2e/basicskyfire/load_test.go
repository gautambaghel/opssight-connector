/*
Copyright (C) 2018 Black Duck Software, Inc.

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

package basicskyfire

import (
	"fmt"
	"testing"
	"time"

	skyfire "github.com/blackducksoftware/perceptor-skyfire/pkg/report"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLoad(t *testing.T) {
	skyfireURL := fmt.Sprintf("http://%s:%s/latestreport", skyfireHost, skyfirePort)
	LoadTests(skyfireURL)
	RegisterFailHandler(Fail)
	RunSpecs(t, "load-test")
}

func LoadTests(skyfireURL string) {
	fmt.Printf("skyfireURL: %s\n", skyfireURL)
	var report *skyfire.Report
	var err error
	for {
		report, err = fetchSkyfireReport(skyfireURL)
		if err != nil {
			Fail(fmt.Sprintf("unable to fetch skyfire report from %s: %s", skyfireURL, err.Error()))
			return
		}

		fmt.Printf("report: %v", report)
		if report != nil {
			break
		} else {
			time.Sleep(10 * time.Second)
		}
	}

	fmt.Println("Outside the for loop")

	dockerClient, err := NewDocker()
	if err != nil {
		fmt.Errorf("Unable to instantiate Docker client due to %+v", err)
	}
	images := dockerClient.GetDockerImages(noOfPods)

	for _, image := range images {
		fmt.Printf("pod name: %s, image: %s:%s \n", image.PodName, image.ImageName, image.Tag)
		addPods(image.PodName, fmt.Sprintf("%s:%s", image.ImageName, image.Tag), int32(3007))
	}

	createPods(configPath)

	// TODO: write test cases to verify the created pod by using skyfire
}