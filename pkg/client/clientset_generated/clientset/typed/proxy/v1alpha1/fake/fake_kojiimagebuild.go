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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/vfreex/release-apiserver/pkg/apis/proxy/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKojiImageBuilds implements KojiImageBuildInterface
type FakeKojiImageBuilds struct {
	Fake *FakeProxyV1alpha1
	ns   string
}

var kojiimagebuildsResource = schema.GroupVersionResource{Group: "proxy.art.openshift.io", Version: "v1alpha1", Resource: "kojiimagebuilds"}

var kojiimagebuildsKind = schema.GroupVersionKind{Group: "proxy.art.openshift.io", Version: "v1alpha1", Kind: "KojiImageBuild"}

// Get takes name of the kojiImageBuild, and returns the corresponding kojiImageBuild object, and an error if there is any.
func (c *FakeKojiImageBuilds) Get(name string, options v1.GetOptions) (result *v1alpha1.KojiImageBuild, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kojiimagebuildsResource, c.ns, name), &v1alpha1.KojiImageBuild{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KojiImageBuild), err
}

// List takes label and field selectors, and returns the list of KojiImageBuilds that match those selectors.
func (c *FakeKojiImageBuilds) List(opts v1.ListOptions) (result *v1alpha1.KojiImageBuildList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kojiimagebuildsResource, kojiimagebuildsKind, c.ns, opts), &v1alpha1.KojiImageBuildList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KojiImageBuildList{ListMeta: obj.(*v1alpha1.KojiImageBuildList).ListMeta}
	for _, item := range obj.(*v1alpha1.KojiImageBuildList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kojiImageBuilds.
func (c *FakeKojiImageBuilds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kojiimagebuildsResource, c.ns, opts))

}

// Create takes the representation of a kojiImageBuild and creates it.  Returns the server's representation of the kojiImageBuild, and an error, if there is any.
func (c *FakeKojiImageBuilds) Create(kojiImageBuild *v1alpha1.KojiImageBuild) (result *v1alpha1.KojiImageBuild, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kojiimagebuildsResource, c.ns, kojiImageBuild), &v1alpha1.KojiImageBuild{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KojiImageBuild), err
}

// Update takes the representation of a kojiImageBuild and updates it. Returns the server's representation of the kojiImageBuild, and an error, if there is any.
func (c *FakeKojiImageBuilds) Update(kojiImageBuild *v1alpha1.KojiImageBuild) (result *v1alpha1.KojiImageBuild, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kojiimagebuildsResource, c.ns, kojiImageBuild), &v1alpha1.KojiImageBuild{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KojiImageBuild), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKojiImageBuilds) UpdateStatus(kojiImageBuild *v1alpha1.KojiImageBuild) (*v1alpha1.KojiImageBuild, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kojiimagebuildsResource, "status", c.ns, kojiImageBuild), &v1alpha1.KojiImageBuild{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KojiImageBuild), err
}

// Delete takes name of the kojiImageBuild and deletes it. Returns an error if one occurs.
func (c *FakeKojiImageBuilds) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kojiimagebuildsResource, c.ns, name), &v1alpha1.KojiImageBuild{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKojiImageBuilds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kojiimagebuildsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.KojiImageBuildList{})
	return err
}

// Patch applies the patch and returns the patched kojiImageBuild.
func (c *FakeKojiImageBuilds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KojiImageBuild, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kojiimagebuildsResource, c.ns, name, pt, data, subresources...), &v1alpha1.KojiImageBuild{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KojiImageBuild), err
}