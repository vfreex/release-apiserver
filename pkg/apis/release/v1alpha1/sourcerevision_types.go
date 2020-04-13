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

// SourceRevision
// +k8s:openapi-gen=true
// +resource:path=sourcerevisions,shortname=sr,strategy=SourceRevisionStrategy
type SourceRevision struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SourceRevisionSpec   `json:"spec,omitempty"`
	Status SourceRevisionStatus `json:"status,omitempty"`
}

// SourceRevisionSpec defines the desired state of SourceRevision
type SourceRevisionSpec struct {
	ComponentName string `json:"componentName"`
	CommitHash    string `json:"commitHash,omitempty"`
}

// SourceRevisionStatus defines the observed state of SourceRevision
type SourceRevisionStatus struct {
}
