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

// Advisory
// +k8s:openapi-gen=true
// +resource:path=advisories,shortname=ad,strategy=AdvisoryStrategy
type Advisory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AdvisorySpec   `json:"spec,omitempty"`
	Status AdvisoryStatus `json:"status,omitempty"`
}

// AdvisorySpec defines the desired state of Advisory
type AdvisorySpec struct {
	Number      int    `json:"number,omitempty"`
	Impetus     string `json:"impetus,omitempty"`
	ReleaseName string `json:"releaseName,omitempty"`
	Instance    string `json:"instance,omitempty"`
}

// AdvisoryStatus defines the observed state of Advisory
type AdvisoryStatus struct {
	LiveID string `json:"liveID,omitempty"`
	State  string `json:"state,omitempty"`
}
