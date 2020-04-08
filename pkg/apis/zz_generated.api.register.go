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

package apis

import (
	"github.com/vfreex/release-apiserver/pkg/apis/release"
	_ "github.com/vfreex/release-apiserver/pkg/apis/release/install" // Install the release group
	releasev1alpha1 "github.com/vfreex/release-apiserver/pkg/apis/release/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/apiserver-builder-alpha/pkg/builders"
)

var (
	localSchemeBuilder = runtime.SchemeBuilder{
		releasev1alpha1.AddToScheme,
	}
	AddToScheme = localSchemeBuilder.AddToScheme
)

// GetAllApiBuilders returns all known APIGroupBuilders
// so they can be registered with the apiserver
func GetAllApiBuilders() []*builders.APIGroupBuilder {
	return []*builders.APIGroupBuilder{
		GetReleaseAPIBuilder(),
	}
}

var releaseApiGroup = builders.NewApiGroupBuilder(
	"release.art.openshift.io",
	"github.com/vfreex/release-apiserver/pkg/apis/release").
	WithUnVersionedApi(release.ApiVersion).
	WithVersionedApis(
		releasev1alpha1.ApiVersion,
	).
	WithRootScopedKinds()

func GetReleaseAPIBuilder() *builders.APIGroupBuilder {
	return releaseApiGroup
}
