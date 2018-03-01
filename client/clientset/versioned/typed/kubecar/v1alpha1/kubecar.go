/*
Copyright 2018 The Voyager Authors.

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
package v1alpha1

import (
	v1alpha1 "github.com/emruz-hossain/k8s-admission-webhook-with-extension-apiserver/apis/kubecar/v1alpha1"
	scheme "github.com/emruz-hossain/k8s-admission-webhook-with-extension-apiserver/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KubecarsGetter has a method to return a KubecarInterface.
// A group's client should implement this interface.
type KubecarsGetter interface {
	Kubecars(namespace string) KubecarInterface
}

// KubecarInterface has methods to work with Kubecar resources.
type KubecarInterface interface {
	Create(*v1alpha1.Kubecar) (*v1alpha1.Kubecar, error)
	Update(*v1alpha1.Kubecar) (*v1alpha1.Kubecar, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Kubecar, error)
	List(opts v1.ListOptions) (*v1alpha1.KubecarList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Kubecar, err error)
	KubecarExpansion
}

// kubecars implements KubecarInterface
type kubecars struct {
	client rest.Interface
	ns     string
}

// newKubecars returns a Kubecars
func newKubecars(c *KubecarV1alpha1Client, namespace string) *kubecars {
	return &kubecars{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kubecar, and returns the corresponding kubecar object, and an error if there is any.
func (c *kubecars) Get(name string, options v1.GetOptions) (result *v1alpha1.Kubecar, err error) {
	result = &v1alpha1.Kubecar{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kubecars").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Kubecars that match those selectors.
func (c *kubecars) List(opts v1.ListOptions) (result *v1alpha1.KubecarList, err error) {
	result = &v1alpha1.KubecarList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kubecars").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kubecars.
func (c *kubecars) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kubecars").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a kubecar and creates it.  Returns the server's representation of the kubecar, and an error, if there is any.
func (c *kubecars) Create(kubecar *v1alpha1.Kubecar) (result *v1alpha1.Kubecar, err error) {
	result = &v1alpha1.Kubecar{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kubecars").
		Body(kubecar).
		Do().
		Into(result)
	return
}

// Update takes the representation of a kubecar and updates it. Returns the server's representation of the kubecar, and an error, if there is any.
func (c *kubecars) Update(kubecar *v1alpha1.Kubecar) (result *v1alpha1.Kubecar, err error) {
	result = &v1alpha1.Kubecar{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kubecars").
		Name(kubecar.Name).
		Body(kubecar).
		Do().
		Into(result)
	return
}

// Delete takes name of the kubecar and deletes it. Returns an error if one occurs.
func (c *kubecars) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kubecars").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kubecars) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kubecars").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kubecar.
func (c *kubecars) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Kubecar, err error) {
	result = &v1alpha1.Kubecar{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kubecars").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
