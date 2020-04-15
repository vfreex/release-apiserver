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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/vfreex/release-apiserver/pkg/apis/art/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ComponentLister helps list Components.
type ComponentLister interface {
	// List lists all Components in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Component, err error)
	// Components returns an object that can list and get Components.
	Components(namespace string) ComponentNamespaceLister
	ComponentListerExpansion
}

// componentLister implements the ComponentLister interface.
type componentLister struct {
	indexer cache.Indexer
}

// NewComponentLister returns a new ComponentLister.
func NewComponentLister(indexer cache.Indexer) ComponentLister {
	return &componentLister{indexer: indexer}
}

// List lists all Components in the indexer.
func (s *componentLister) List(selector labels.Selector) (ret []*v1alpha1.Component, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Component))
	})
	return ret, err
}

// Components returns an object that can list and get Components.
func (s *componentLister) Components(namespace string) ComponentNamespaceLister {
	return componentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ComponentNamespaceLister helps list and get Components.
type ComponentNamespaceLister interface {
	// List lists all Components in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Component, err error)
	// Get retrieves the Component from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Component, error)
	ComponentNamespaceListerExpansion
}

// componentNamespaceLister implements the ComponentNamespaceLister
// interface.
type componentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Components in the indexer for a given namespace.
func (s componentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Component, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Component))
	})
	return ret, err
}

// Get retrieves the Component from the indexer for a given namespace and name.
func (s componentNamespaceLister) Get(name string) (*v1alpha1.Component, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("component"), name)
	}
	return obj.(*v1alpha1.Component), nil
}
