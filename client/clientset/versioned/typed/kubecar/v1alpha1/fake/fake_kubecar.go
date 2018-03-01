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
package fake

import (
	v1alpha1 "k8s-admission-webhook-with-extension-apiserver/apis/kubecar/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKubecars implements KubecarInterface
type FakeKubecars struct {
	Fake *FakeKubecarV1alpha1
	ns   string
}

var kubecarsResource = schema.GroupVersionResource{Group: "kubecar.emruz.com", Version: "v1alpha1", Resource: "kubecars"}

var kubecarsKind = schema.GroupVersionKind{Group: "kubecar.emruz.com", Version: "v1alpha1", Kind: "Kubecar"}

// Get takes name of the kubecar, and returns the corresponding kubecar object, and an error if there is any.
func (c *FakeKubecars) Get(name string, options v1.GetOptions) (result *v1alpha1.Kubecar, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kubecarsResource, c.ns, name), &v1alpha1.Kubecar{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Kubecar), err
}

// List takes label and field selectors, and returns the list of Kubecars that match those selectors.
func (c *FakeKubecars) List(opts v1.ListOptions) (result *v1alpha1.KubecarList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kubecarsResource, kubecarsKind, c.ns, opts), &v1alpha1.KubecarList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KubecarList{}
	for _, item := range obj.(*v1alpha1.KubecarList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kubecars.
func (c *FakeKubecars) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kubecarsResource, c.ns, opts))

}

// Create takes the representation of a kubecar and creates it.  Returns the server's representation of the kubecar, and an error, if there is any.
func (c *FakeKubecars) Create(kubecar *v1alpha1.Kubecar) (result *v1alpha1.Kubecar, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kubecarsResource, c.ns, kubecar), &v1alpha1.Kubecar{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Kubecar), err
}

// Update takes the representation of a kubecar and updates it. Returns the server's representation of the kubecar, and an error, if there is any.
func (c *FakeKubecars) Update(kubecar *v1alpha1.Kubecar) (result *v1alpha1.Kubecar, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kubecarsResource, c.ns, kubecar), &v1alpha1.Kubecar{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Kubecar), err
}

// Delete takes name of the kubecar and deletes it. Returns an error if one occurs.
func (c *FakeKubecars) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kubecarsResource, c.ns, name), &v1alpha1.Kubecar{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKubecars) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kubecarsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.KubecarList{})
	return err
}

// Patch applies the patch and returns the patched kubecar.
func (c *FakeKubecars) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Kubecar, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kubecarsResource, c.ns, name, data, subresources...), &v1alpha1.Kubecar{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Kubecar), err
}
