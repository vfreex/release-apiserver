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
// +resource:path=releasestreams,strategy=ReleaseStreamStrategy,shortNames=rs,categories=all;art
type ReleaseStream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   ReleaseStreamSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status ReleaseStreamStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// ReleaseStreamSpec defines the desired state of ReleaseStream
type ReleaseStreamSpec struct {
	OcpBuildData ReleaseStreamOcpBuildData `json:"ocpBuildData,omitempty" protobuf:"bytes,1,opt,name=ocpBuildData"`
}

// ReleaseStreamStatus defines the observed state of ReleaseStream
type ReleaseStreamStatus struct {
}

type ReleaseStreamOcpBuildData struct {
	Git OcpBuildDataGitSource `json:"git,omitempty" protobuf:"bytes,1,opt,name=git"`
}

type OcpBuildDataGitSource struct {
	Url  string `json:"url" protobuf:"bytes,1,opt,name=url"`
	Ref  string `json:"ref" protobuf:"bytes,2,opt,name=ref"`
	Path string `json:"path,omitempty" protobuf:"bytes,3,opt,name=path"`
}
