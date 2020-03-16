// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package proxy

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KojiImageBuild) DeepCopyInto(out *KojiImageBuild) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KojiImageBuild.
func (in *KojiImageBuild) DeepCopy() *KojiImageBuild {
	if in == nil {
		return nil
	}
	out := new(KojiImageBuild)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KojiImageBuild) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KojiImageBuildList) DeepCopyInto(out *KojiImageBuildList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KojiImageBuild, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KojiImageBuildList.
func (in *KojiImageBuildList) DeepCopy() *KojiImageBuildList {
	if in == nil {
		return nil
	}
	out := new(KojiImageBuildList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KojiImageBuildList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KojiImageBuildSpec) DeepCopyInto(out *KojiImageBuildSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KojiImageBuildSpec.
func (in *KojiImageBuildSpec) DeepCopy() *KojiImageBuildSpec {
	if in == nil {
		return nil
	}
	out := new(KojiImageBuildSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KojiImageBuildStatus) DeepCopyInto(out *KojiImageBuildStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KojiImageBuildStatus.
func (in *KojiImageBuildStatus) DeepCopy() *KojiImageBuildStatus {
	if in == nil {
		return nil
	}
	out := new(KojiImageBuildStatus)
	in.DeepCopyInto(out)
	return out
}