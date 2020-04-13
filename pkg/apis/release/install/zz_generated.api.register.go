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

package install

import (
	"github.com/vfreex/release-apiserver/pkg/apis/release"
	"github.com/vfreex/release-apiserver/pkg/apis/release/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"sigs.k8s.io/apiserver-builder-alpha/pkg/builders"
)

func init() {
	Install(builders.Scheme)
}

func Install(scheme *runtime.Scheme) {
	utilruntime.Must(v1alpha1.AddToScheme(scheme))
	utilruntime.Must(release.AddToScheme(scheme))
	utilruntime.Must(addKnownTypes(scheme))
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(release.SchemeGroupVersion,
		&release.Advisory{},
		&release.AdvisoryList{},
		&release.Build{},
		&release.BuildList{},
		&release.Component{},
		&release.ComponentList{},
		&release.Payload{},
		&release.PayloadList{},
		&release.Release{},
		&release.ReleaseList{},
		&release.ReleaseStream{},
		&release.ReleaseStreamList{},
		&release.SourceRevision{},
		&release.SourceRevisionList{},
	)
	return nil
}
