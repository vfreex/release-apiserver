
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


package kojiimagebuildadmission

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

var _ admission.Interface 											= &kojiimagebuildPlugin{}
var _ admission.MutationInterface 									= &kojiimagebuildPlugin{}
var _ admission.ValidationInterface 								= &kojiimagebuildPlugin{}
var _ genericadmissioninitializer.WantsExternalKubeInformerFactory 	= &kojiimagebuildPlugin{}
var _ genericadmissioninitializer.WantsExternalKubeClientSet 		= &kojiimagebuildPlugin{}
var _ aggregatedadmission.WantsAggregatedResourceInformerFactory 	= &kojiimagebuildPlugin{}
var _ aggregatedadmission.WantsAggregatedResourceClientSet 			= &kojiimagebuildPlugin{}

func NewKojiImageBuildPlugin() *kojiimagebuildPlugin {
	return &kojiimagebuildPlugin{
		Handler: admission.NewHandler(admission.Create, admission.Update),
	}
}

type kojiimagebuildPlugin struct {
	*admission.Handler
}

func (p *kojiimagebuildPlugin) ValidateInitialization() error {
	return nil
}

func (p *kojiimagebuildPlugin) Admit(ctx context.Context, a admission.Attributes, o admission.ObjectInterfaces) error {
	return nil
}

func (p *kojiimagebuildPlugin) Validate(ctx context.Context, a admission.Attributes, o admission.ObjectInterfaces) error {
	return nil
}

func (p *kojiimagebuildPlugin) SetAggregatedResourceInformerFactory(aggregatedinformerfactory.SharedInformerFactory) {}

func (p *kojiimagebuildPlugin) SetAggregatedResourceClientSet(aggregatedclientset.Interface) {}

func (p *kojiimagebuildPlugin) SetExternalKubeInformerFactory(informers.SharedInformerFactory) {}

func (p *kojiimagebuildPlugin) SetExternalKubeClientSet(kubernetes.Interface) {}
