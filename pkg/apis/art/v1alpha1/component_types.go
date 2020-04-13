
/*
Copyright 2020 The OpenShift Release APIServer Authors.

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

// Component
// +k8s:openapi-gen=true
// +resource:path=components,strategy=ComponentStrategy
type Component struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComponentSpec   `json:"spec,omitempty"`
	Status ComponentStatus `json:"status,omitempty"`
}


// ComponentSpec defines the desired state of Component
type ComponentSpec struct {
	Source  ComponentSource  `json:"source,omitempty"`
	DistGit ComponentDistGit `json:"distGit,omitempty"`
}

type ComponentSource struct {
	Git ComponentGitSource `json:"git,omitempty"`
}

type ComponentGitSource struct {
	Url  string `json:"url"`
	Ref  string `json:"ref"`
	Path string `json:"path,omitempty"`
}

type ComponentDistGit struct {
	Namespace  string `json:"namespace"`
	Repository string `json:"repository"`
	Instance   string `json:"instance,omitempty"`
}

// ComponentStatus defines the observed state of Component
type ComponentStatus struct {
}
