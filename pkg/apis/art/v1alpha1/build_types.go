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
// +resource:path=builds,strategy=BuildStrategy,shortNames=b,categories=all;art
type Build struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   BuildSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status BuildStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// BuildSpec defines the desired state of Build
type BuildSpec struct {
	Type    string           `json:"type" protobuf:"bytes,1,opt,name=type"`
	Brew    BrewBuild        `json:"brew" protobuf:"bytes,2,opt,name=brew"`
	Jenkins BuildJenkinsInfo `json:"jenkins,omitempty" protobuf:"bytes,3,opt,name=jenkins"`
}

type BrewBuild struct {
	Name       string `json:"name" protobuf:"bytes,1,opt,name=name"`
	Version    string `json:"version" protobuf:"bytes,2,opt,name=version"`
	Release    string `json:"release" protobuf:"bytes,3,opt,name=release"`
	NVR        string `json:"nvr,omitempty" protobuf:"bytes,4,opt,name=nvr"`
	Target     string `json:"target,omitempty" protobuf:"bytes,5,opt,name=target"`
	TaskNumber int32  `json:"taskNumber,omitempty" protobuf:"varint,6,opt,name=taskNumber"`
	Instance   string `json:"instance,omitempty" protobuf:"bytes,7,opt,name=instance"`
}

type BuildJenkinsInfo struct {
	BuildUrl string `json:"buildUrl,omitempty" protobuf:"bytes,1,opt,name=buildUrl"`
}

// BuildStatus defines the observed state of Build
type BuildStatus struct {
}
