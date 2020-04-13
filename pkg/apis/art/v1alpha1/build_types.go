
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

// Build
// +k8s:openapi-gen=true
// +resource:path=builds,strategy=BuildStrategy
type Build struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BuildSpec   `json:"spec,omitempty"`
	Status BuildStatus `json:"status,omitempty"`
}


// BuildSpec defines the desired state of Build
type BuildSpec struct {
	Type    string           `json:"type"`
	Brew    BrewBuild        `json:"brew"`
	Jenkins BuildJenkinsInfo `json:"jenkins,omitempty"`
}

type BrewBuild struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	Release    string `json:"release"`
	NVR        string `json:"nvr,omitempty"`
	Target     string `json:"target,omitempty"`
	TaskNumber int    `json:"taskNumber,omitempty"`
	Instance   string `json:"instance,omitempty"`
}

type BuildJenkinsInfo struct {
	BuildUrl string `json:"buildUrl,omitempty"`
}

// BuildStatus defines the observed state of Build
type BuildStatus struct {
}
