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

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/vfreex/release-apiserver/pkg/apis/art/v1alpha1"
	scheme "github.com/vfreex/release-apiserver/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AdvisoriesGetter has a method to return a AdvisoryInterface.
// A group's client should implement this interface.
type AdvisoriesGetter interface {
	Advisories(namespace string) AdvisoryInterface
}

// AdvisoryInterface has methods to work with Advisory resources.
type AdvisoryInterface interface {
	Create(*v1alpha1.Advisory) (*v1alpha1.Advisory, error)
	Update(*v1alpha1.Advisory) (*v1alpha1.Advisory, error)
	UpdateStatus(*v1alpha1.Advisory) (*v1alpha1.Advisory, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Advisory, error)
	List(opts v1.ListOptions) (*v1alpha1.AdvisoryList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Advisory, err error)
	AdvisoryExpansion
}

// advisories implements AdvisoryInterface
type advisories struct {
	client rest.Interface
	ns     string
}

// newAdvisories returns a Advisories
func newAdvisories(c *ArtV1alpha1Client, namespace string) *advisories {
	return &advisories{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the advisory, and returns the corresponding advisory object, and an error if there is any.
func (c *advisories) Get(name string, options v1.GetOptions) (result *v1alpha1.Advisory, err error) {
	result = &v1alpha1.Advisory{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("advisories").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Advisories that match those selectors.
func (c *advisories) List(opts v1.ListOptions) (result *v1alpha1.AdvisoryList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AdvisoryList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("advisories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested advisories.
func (c *advisories) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("advisories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a advisory and creates it.  Returns the server's representation of the advisory, and an error, if there is any.
func (c *advisories) Create(advisory *v1alpha1.Advisory) (result *v1alpha1.Advisory, err error) {
	result = &v1alpha1.Advisory{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("advisories").
		Body(advisory).
		Do().
		Into(result)
	return
}

// Update takes the representation of a advisory and updates it. Returns the server's representation of the advisory, and an error, if there is any.
func (c *advisories) Update(advisory *v1alpha1.Advisory) (result *v1alpha1.Advisory, err error) {
	result = &v1alpha1.Advisory{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("advisories").
		Name(advisory.Name).
		Body(advisory).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *advisories) UpdateStatus(advisory *v1alpha1.Advisory) (result *v1alpha1.Advisory, err error) {
	result = &v1alpha1.Advisory{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("advisories").
		Name(advisory.Name).
		SubResource("status").
		Body(advisory).
		Do().
		Into(result)
	return
}

// Delete takes name of the advisory and deletes it. Returns an error if one occurs.
func (c *advisories) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("advisories").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *advisories) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("advisories").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched advisory.
func (c *advisories) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Advisory, err error) {
	result = &v1alpha1.Advisory{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("advisories").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}