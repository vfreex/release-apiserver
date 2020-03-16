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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KojiImageBuild
// +k8s:openapi-gen=true
// +resource:path=kojiimagebuilds,strategy=KojiImageBuildStrategy,shortname=ki,rest=KojiImageBuildREST
type KojiImageBuild struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KojiImageBuildSpec   `json:"spec,omitempty"`
	Status KojiImageBuildStatus `json:"status,omitempty"`
}

// KojiImageBuildSpec defines the desired state of KojiImageBuild
type KojiImageBuildSpec struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	Release string `json:"release,omitempty"`
}

// KojiImageBuildStatus defines the observed state of KojiImageBuild
type KojiImageBuildStatus struct {
	KojiState     string `json:"kojiState,omitempty"`
	KojiBuildLink string `json:"kojiBuildLink,omitempty"`
}
