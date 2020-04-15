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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	artv1alpha1 "github.com/vfreex/release-apiserver/pkg/apis/art/v1alpha1"
	clientset "github.com/vfreex/release-apiserver/pkg/client/clientset_generated/clientset"
	internalinterfaces "github.com/vfreex/release-apiserver/pkg/client/informers_generated/externalversions/internalinterfaces"
	v1alpha1 "github.com/vfreex/release-apiserver/pkg/client/listers_generated/art/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AdvisoryInformer provides access to a shared informer and lister for
// Advisories.
type AdvisoryInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.AdvisoryLister
}

type advisoryInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAdvisoryInformer constructs a new informer for Advisory type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAdvisoryInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAdvisoryInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAdvisoryInformer constructs a new informer for Advisory type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAdvisoryInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ArtV1alpha1().Advisories(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ArtV1alpha1().Advisories(namespace).Watch(options)
			},
		},
		&artv1alpha1.Advisory{},
		resyncPeriod,
		indexers,
	)
}

func (f *advisoryInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAdvisoryInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *advisoryInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&artv1alpha1.Advisory{}, f.defaultInformer)
}

func (f *advisoryInformer) Lister() v1alpha1.AdvisoryLister {
	return v1alpha1.NewAdvisoryLister(f.Informer().GetIndexer())
}
