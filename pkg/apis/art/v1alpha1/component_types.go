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
// +resource:path=components,strategy=ComponentStrategy,shortNames=cp,categories=all;art
type Component struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   ComponentSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status ComponentStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// ComponentSpec defines the desired state of Component
type ComponentSpec struct {
	Kind    string           `json:"kind,omitempty" protobuf:"bytes,1,opt,name=kind"`
	Image   ComponentImage   `json:"image,omitempty" protobuf:"bytes,2,opt,name=image"`
	Source  ComponentSource  `json:"source,omitempty" protobuf:"bytes,3,opt,name=source"`
	DistGit ComponentDistGit `json:"distGit,omitempty" protobuf:"bytes,4,opt,name=distGit"`
}

type ComponentImage struct {
	Repository string `json:"repository" protobuf:"bytes,1,opt,name=repository"`
	Namespace  string `json:"namespace" protobuf:"bytes,2,opt,name=namespace"`
}

type ComponentSource struct {
	Git ComponentGitSource `json:"git,omitempty" protobuf:"bytes,1,opt,name=git"`
}

type ComponentGitSource struct {
	Url         string `json:"url" protobuf:"bytes,1,opt,name=url"`
	Ref         string `json:"ref" protobuf:"bytes,2,opt,name=ref"`
	FallbackRef string `json:"fallbackRef,omitempty" protobuf:"bytes,3,opt,name=fallbackRef"`
	Path        string `json:"path,omitempty" protobuf:"bytes,4,opt,name=path"`
}

type ComponentDistGit struct {
	Namespace  string `json:"namespace" protobuf:"bytes,1,opt,name=namespace"`
	Repository string `json:"repository" protobuf:"bytes,2,opt,name=repository"`
	Instance   string `json:"instance,omitempty" protobuf:"bytes,3,opt,name=instance"`
}

// ComponentStatus defines the observed state of Component
type ComponentStatus struct {
}
