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

// Code generated by apiregister-gen. DO NOT EDIT.

package release

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/rest"
	"sigs.k8s.io/apiserver-builder-alpha/pkg/builders"
)

var (
	ReleaseReleaseStreamStorage = builders.NewApiResource( // Resource status endpoint
		InternalReleaseStream,
		func() runtime.Object { return &ReleaseStream{} },     // Register versioned resource
		func() runtime.Object { return &ReleaseStreamList{} }, // Register versioned resource list
		&ReleaseStreamStrategy{builders.StorageStrategySingleton},
	)
	InternalReleaseStream = builders.NewInternalResource(
		"releasestreams",
		"ReleaseStream",
		func() runtime.Object { return &ReleaseStream{} },
		func() runtime.Object { return &ReleaseStreamList{} },
	)
	InternalReleaseStreamStatus = builders.NewInternalResourceStatus(
		"releasestreams",
		"ReleaseStreamStatus",
		func() runtime.Object { return &ReleaseStream{} },
		func() runtime.Object { return &ReleaseStreamList{} },
	)
	// Registered resources and subresources
	ApiVersion = builders.NewApiGroup("release.art.openshift.io").WithKinds(
		InternalReleaseStream,
		InternalReleaseStreamStatus,
	)

	// Required by code generated by go2idl
	AddToScheme = (&runtime.SchemeBuilder{
		ApiVersion.SchemeBuilder.AddToScheme,
		RegisterDefaults,
	}).AddToScheme
	SchemeBuilder      = ApiVersion.SchemeBuilder
	localSchemeBuilder = &SchemeBuilder
	SchemeGroupVersion = ApiVersion.GroupVersion
)

// Required by code generated by go2idl
// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Required by code generated by go2idl
// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// +genclient
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ReleaseStream struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   ReleaseStreamSpec
	Status ReleaseStreamStatus
}

type ReleaseStreamGitSource struct {
	Url  string
	Ref  string
	Path string
}

type ReleaseStreamOcpBuildData struct {
	Git ReleaseStreamGitSource
}

type ReleaseStreamSpec struct {
	OcpBuildData ReleaseStreamOcpBuildData
}

type ReleaseStreamStatus struct {
}

//
// ReleaseStream Functions and Structs
//
// +k8s:deepcopy-gen=false
type ReleaseStreamStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type ReleaseStreamStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ReleaseStreamList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []ReleaseStream
}

func (ReleaseStream) NewStatus() interface{} {
	return ReleaseStreamStatus{}
}

func (pc *ReleaseStream) GetStatus() interface{} {
	return pc.Status
}

func (pc *ReleaseStream) SetStatus(s interface{}) {
	pc.Status = s.(ReleaseStreamStatus)
}

func (pc *ReleaseStream) GetSpec() interface{} {
	return pc.Spec
}

func (pc *ReleaseStream) SetSpec(s interface{}) {
	pc.Spec = s.(ReleaseStreamSpec)
}

func (pc *ReleaseStream) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *ReleaseStream) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc ReleaseStream) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store ReleaseStream.
// +k8s:deepcopy-gen=false
type ReleaseStreamRegistry interface {
	ListReleaseStreams(ctx context.Context, options *internalversion.ListOptions) (*ReleaseStreamList, error)
	GetReleaseStream(ctx context.Context, id string, options *metav1.GetOptions) (*ReleaseStream, error)
	CreateReleaseStream(ctx context.Context, id *ReleaseStream) (*ReleaseStream, error)
	UpdateReleaseStream(ctx context.Context, id *ReleaseStream) (*ReleaseStream, error)
	DeleteReleaseStream(ctx context.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewReleaseStreamRegistry(sp builders.StandardStorageProvider) ReleaseStreamRegistry {
	return &storageReleaseStream{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageReleaseStream struct {
	builders.StandardStorageProvider
}

func (s *storageReleaseStream) ListReleaseStreams(ctx context.Context, options *internalversion.ListOptions) (*ReleaseStreamList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*ReleaseStreamList), err
}

func (s *storageReleaseStream) GetReleaseStream(ctx context.Context, id string, options *metav1.GetOptions) (*ReleaseStream, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*ReleaseStream), nil
}

func (s *storageReleaseStream) CreateReleaseStream(ctx context.Context, object *ReleaseStream) (*ReleaseStream, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, &metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*ReleaseStream), nil
}

func (s *storageReleaseStream) UpdateReleaseStream(ctx context.Context, object *ReleaseStream) (*ReleaseStream, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil, false, &metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*ReleaseStream), nil
}

func (s *storageReleaseStream) DeleteReleaseStream(ctx context.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil, &metav1.DeleteOptions{})
	return sync, err
}
