
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


package main

import (
	_ "github.com/go-openapi/loads"
	// Make sure dep tools picks up these dependencies
	_ "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/apiserver-builder-alpha/pkg/apiserver"

	_ "k8s.io/client-go/plugin/pkg/client/auth" // Enable cloud provider auth
	"sigs.k8s.io/apiserver-builder-alpha/pkg/cmd/server"

	"github.com/vfreex/release-apiserver/pkg/apis"
	"github.com/vfreex/release-apiserver/pkg/openapi"
	_ "github.com/vfreex/release-apiserver/plugin/admission/install"
	"k8s.io/apiserver/pkg/authorization/authorizerfactory"
)

func main() {
	version := "v0"

	err := server.StartApiServerWithOptions(&server.StartOptions{
		EtcdPath:         "/registry/openshift.io",
		Apis:             apis.GetAllApiBuilders(),
		Openapidefs:      openapi.GetOpenAPIDefinitions,
		Title:            "Api",
		Version:          version,

		TweakConfigFuncs: []func(config *apiserver.Config) error {
			func(config *apiserver.Config) error {
				// Currently we only permit system:masters and art:apiclients groups to use this apiserver
				config.RecommendedConfig.Authorization.Authorizer = authorizerfactory.NewPrivilegedGroups("system:masters", "art:apiclients")
				return nil
			},
		},
		// FlagConfigFuncs []func(*cobra.Command) error
	})
	if err != nil {
		panic(err)
	}
}
