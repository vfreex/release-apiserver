
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


package releasestreamadmission

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

var _ admission.Interface 											= &releasestreamPlugin{}
var _ admission.MutationInterface 									= &releasestreamPlugin{}
var _ admission.ValidationInterface 								= &releasestreamPlugin{}
var _ genericadmissioninitializer.WantsExternalKubeInformerFactory 	= &releasestreamPlugin{}
var _ genericadmissioninitializer.WantsExternalKubeClientSet 		= &releasestreamPlugin{}
var _ aggregatedadmission.WantsAggregatedResourceInformerFactory 	= &releasestreamPlugin{}
var _ aggregatedadmission.WantsAggregatedResourceClientSet 			= &releasestreamPlugin{}

func NewReleaseStreamPlugin() *releasestreamPlugin {
	return &releasestreamPlugin{
		Handler: admission.NewHandler(admission.Create, admission.Update),
	}
}

type releasestreamPlugin struct {
	*admission.Handler
}

func (p *releasestreamPlugin) ValidateInitialization() error {
	return nil
}

func (p *releasestreamPlugin) Admit(ctx context.Context, a admission.Attributes, o admission.ObjectInterfaces) error {
	return nil
}

func (p *releasestreamPlugin) Validate(ctx context.Context, a admission.Attributes, o admission.ObjectInterfaces) error {
	return nil
}

func (p *releasestreamPlugin) SetAggregatedResourceInformerFactory(aggregatedinformerfactory.SharedInformerFactory) {}

func (p *releasestreamPlugin) SetAggregatedResourceClientSet(aggregatedclientset.Interface) {}

func (p *releasestreamPlugin) SetExternalKubeInformerFactory(informers.SharedInformerFactory) {}

func (p *releasestreamPlugin) SetExternalKubeClientSet(kubernetes.Interface) {}
