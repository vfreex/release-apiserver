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

// RPMCompose
// +k8s:openapi-gen=true
// +resource:path=rpmcomposes,strategy=RPMComposeStrategy,shortNames=rc,categories=all;art
type RPMCompose struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   RPMComposeSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status RPMComposeStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// RPMComposeSpec defines the desired state of RPMCompose
type RPMComposeSpec struct {
}

// RPMComposeStatus defines the observed state of RPMCompose
type RPMComposeStatus struct {
}
