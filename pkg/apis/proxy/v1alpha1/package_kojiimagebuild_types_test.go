
/*
Copyright 2020 The OCP Release APIServer Authors.

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


package v1alpha1_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/vfreex/release-apiserver/pkg/apis/proxy/v1alpha1"
	. "github.com/vfreex/release-apiserver/pkg/client/clientset_generated/clientset/typed/proxy/v1alpha1"
)

var _ = Describe("KojiImageBuild", func() {
	var instance KojiImageBuild
	var expected KojiImageBuild
	var client KojiImageBuildInterface

	BeforeEach(func() {
		instance = KojiImageBuild{}
		instance.Name = "instance-1"

		expected = instance
	})

	AfterEach(func() {
		client.Delete(instance.Name, &metav1.DeleteOptions{})
	})

	Describe("when sending a package request", func() {
		It("should return success", func() {
			client = cs.ProxyV1alpha1().Kojiimagebuilds("kojiimagebuild-test-package")
			_, err := client.Create(&instance)
			Expect(err).ShouldNot(HaveOccurred())

			package := &KojiImageBuildPackage{}
			package.Name = instance.Name
			restClient := cs.ProxyV1alpha1().RESTClient()
			err = restClient.Post().Namespace("kojiimagebuild-test-package").
				Name(instance.Name).
				Resource("kojiimagebuilds").
				SubResource("package").
				Body(package).Do().Error()
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
