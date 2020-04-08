
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

// ReleaseStream
// +k8s:openapi-gen=true
// +resource:path=releasestreams,strategy=ReleaseStreamStrategy
type ReleaseStream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ReleaseStreamSpec   `json:"spec,omitempty"`
	Status ReleaseStreamStatus `json:"status,omitempty"`
}

// ReleaseStreamSpec defines the desired state of ReleaseStream
type ReleaseStreamSpec struct {
	OcpBuildData OcpBuildData `json:"ocpBuildData,omitempty"`
}

// ReleaseStreamStatus defines the observed state of ReleaseStream
type ReleaseStreamStatus struct {
}

type OcpBuildData struct {
	Git GitSource `json:"git,omitempty"`
}

type GitSource struct {
	Url string `json:"url,omitempty"`
	Ref string `json:"ref,omitempty"`
	// +kubebuilder:validation:MinLength=1
	Path string `json:"path,omitempty"`
}
