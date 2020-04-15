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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/vfreex/release-apiserver/pkg/apis/art/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeComponents implements ComponentInterface
type FakeComponents struct {
	Fake *FakeArtV1alpha1
	ns   string
}

var componentsResource = schema.GroupVersionResource{Group: "art.openshift.io", Version: "v1alpha1", Resource: "components"}

var componentsKind = schema.GroupVersionKind{Group: "art.openshift.io", Version: "v1alpha1", Kind: "Component"}

// Get takes name of the component, and returns the corresponding component object, and an error if there is any.
func (c *FakeComponents) Get(name string, options v1.GetOptions) (result *v1alpha1.Component, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(componentsResource, c.ns, name), &v1alpha1.Component{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Component), err
}

// List takes label and field selectors, and returns the list of Components that match those selectors.
func (c *FakeComponents) List(opts v1.ListOptions) (result *v1alpha1.ComponentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(componentsResource, componentsKind, c.ns, opts), &v1alpha1.ComponentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComponentList{ListMeta: obj.(*v1alpha1.ComponentList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComponentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested components.
func (c *FakeComponents) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(componentsResource, c.ns, opts))

}

// Create takes the representation of a component and creates it.  Returns the server's representation of the component, and an error, if there is any.
func (c *FakeComponents) Create(component *v1alpha1.Component) (result *v1alpha1.Component, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(componentsResource, c.ns, component), &v1alpha1.Component{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Component), err
}

// Update takes the representation of a component and updates it. Returns the server's representation of the component, and an error, if there is any.
func (c *FakeComponents) Update(component *v1alpha1.Component) (result *v1alpha1.Component, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(componentsResource, c.ns, component), &v1alpha1.Component{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Component), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComponents) UpdateStatus(component *v1alpha1.Component) (*v1alpha1.Component, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(componentsResource, "status", c.ns, component), &v1alpha1.Component{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Component), err
}

// Delete takes name of the component and deletes it. Returns an error if one occurs.
func (c *FakeComponents) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(componentsResource, c.ns, name), &v1alpha1.Component{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComponents) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(componentsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComponentList{})
	return err
}

// Patch applies the patch and returns the patched component.
func (c *FakeComponents) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Component, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(componentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Component{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Component), err
}
