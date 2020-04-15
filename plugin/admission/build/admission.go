
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


package buildadmission

import (
	"context"
	aggregatedadmission "github.com/vfreex/release-apiserver/plugin/admission"
	aggregatedinformerfactory "github.com/vfreex/release-apiserver/pkg/client/informers_generated/externalversions"
	aggregatedclientset "github.com/vfreex/release-apiserver/pkg/client/clientset_generated/clientset"
	genericadmissioninitializer "k8s.io/apiserver/pkg/admission/initializer"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/apiserver/pkg/admission"
)

var _ admission.Interface 											= &buildPlugin{}
var _ admission.MutationInterface 									= &buildPlugin{}
var _ admission.ValidationInterface 								= &buildPlugin{}
var _ genericadmissioninitializer.WantsExternalKubeInformerFactory 	= &buildPlugin{}
var _ genericadmissioninitializer.WantsExternalKubeClientSet 		= &buildPlugin{}
var _ aggregatedadmission.WantsAggregatedResourceInformerFactory 	= &buildPlugin{}
var _ aggregatedadmission.WantsAggregatedResourceClientSet 			= &buildPlugin{}

func NewBuildPlugin() *buildPlugin {
	return &buildPlugin{
		Handler: admission.NewHandler(admission.Create, admission.Update),
	}
}

type buildPlugin struct {
	*admission.Handler
}

func (p *buildPlugin) ValidateInitialization() error {
	return nil
}

func (p *buildPlugin) Admit(ctx context.Context, a admission.Attributes, o admission.ObjectInterfaces) error {
	return nil
}

func (p *buildPlugin) Validate(ctx context.Context, a admission.Attributes, o admission.ObjectInterfaces) error {
	return nil
}

func (p *buildPlugin) SetAggregatedResourceInformerFactory(aggregatedinformerfactory.SharedInformerFactory) {}

func (p *buildPlugin) SetAggregatedResourceClientSet(aggregatedclientset.Interface) {}

func (p *buildPlugin) SetExternalKubeInformerFactory(informers.SharedInformerFactory) {}

func (p *buildPlugin) SetExternalKubeClientSet(kubernetes.Interface) {}
