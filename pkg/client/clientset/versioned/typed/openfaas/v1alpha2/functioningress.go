/*
Copyright 2019 OpenFaaS Authors

Licensed under the MIT license. See LICENSE file in the project root for full license information.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "github.com/openfaas-incubator/ingress-operator/pkg/apis/openfaas/v1alpha2"
	scheme "github.com/openfaas-incubator/ingress-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FunctionIngressesGetter has a method to return a FunctionIngressInterface.
// A group's client should implement this interface.
type FunctionIngressesGetter interface {
	FunctionIngresses(namespace string) FunctionIngressInterface
}

// FunctionIngressInterface has methods to work with FunctionIngress resources.
type FunctionIngressInterface interface {
	Create(*v1alpha2.FunctionIngress) (*v1alpha2.FunctionIngress, error)
	Update(*v1alpha2.FunctionIngress) (*v1alpha2.FunctionIngress, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.FunctionIngress, error)
	List(opts v1.ListOptions) (*v1alpha2.FunctionIngressList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.FunctionIngress, err error)
	FunctionIngressExpansion
}

// functionIngresses implements FunctionIngressInterface
type functionIngresses struct {
	client rest.Interface
	ns     string
}

// newFunctionIngresses returns a FunctionIngresses
func newFunctionIngresses(c *OpenfaasV1alpha2Client, namespace string) *functionIngresses {
	return &functionIngresses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the functionIngress, and returns the corresponding functionIngress object, and an error if there is any.
func (c *functionIngresses) Get(name string, options v1.GetOptions) (result *v1alpha2.FunctionIngress, err error) {
	result = &v1alpha2.FunctionIngress{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("functioningresses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FunctionIngresses that match those selectors.
func (c *functionIngresses) List(opts v1.ListOptions) (result *v1alpha2.FunctionIngressList, err error) {
	result = &v1alpha2.FunctionIngressList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("functioningresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested functionIngresses.
func (c *functionIngresses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("functioningresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a functionIngress and creates it.  Returns the server's representation of the functionIngress, and an error, if there is any.
func (c *functionIngresses) Create(functionIngress *v1alpha2.FunctionIngress) (result *v1alpha2.FunctionIngress, err error) {
	result = &v1alpha2.FunctionIngress{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("functioningresses").
		Body(functionIngress).
		Do().
		Into(result)
	return
}

// Update takes the representation of a functionIngress and updates it. Returns the server's representation of the functionIngress, and an error, if there is any.
func (c *functionIngresses) Update(functionIngress *v1alpha2.FunctionIngress) (result *v1alpha2.FunctionIngress, err error) {
	result = &v1alpha2.FunctionIngress{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("functioningresses").
		Name(functionIngress.Name).
		Body(functionIngress).
		Do().
		Into(result)
	return
}

// Delete takes name of the functionIngress and deletes it. Returns an error if one occurs.
func (c *functionIngresses) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("functioningresses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *functionIngresses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("functioningresses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched functionIngress.
func (c *functionIngresses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.FunctionIngress, err error) {
	result = &v1alpha2.FunctionIngress{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("functioningresses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
