/*
Copyright 2020 The Tekton Authors

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
	"context"

	v1alpha1 "github.com/tektoncd/operator/pkg/apis/operator/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTektonConfigs implements TektonConfigInterface
type FakeTektonConfigs struct {
	Fake *FakeOperatorV1alpha1
}

var tektonconfigsResource = v1alpha1.SchemeGroupVersion.WithResource("tektonconfigs")

var tektonconfigsKind = v1alpha1.SchemeGroupVersion.WithKind("TektonConfig")

// Get takes name of the tektonConfig, and returns the corresponding tektonConfig object, and an error if there is any.
func (c *FakeTektonConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TektonConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(tektonconfigsResource, name), &v1alpha1.TektonConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TektonConfig), err
}

// List takes label and field selectors, and returns the list of TektonConfigs that match those selectors.
func (c *FakeTektonConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TektonConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(tektonconfigsResource, tektonconfigsKind, opts), &v1alpha1.TektonConfigList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TektonConfigList{ListMeta: obj.(*v1alpha1.TektonConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.TektonConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tektonConfigs.
func (c *FakeTektonConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(tektonconfigsResource, opts))
}

// Create takes the representation of a tektonConfig and creates it.  Returns the server's representation of the tektonConfig, and an error, if there is any.
func (c *FakeTektonConfigs) Create(ctx context.Context, tektonConfig *v1alpha1.TektonConfig, opts v1.CreateOptions) (result *v1alpha1.TektonConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(tektonconfigsResource, tektonConfig), &v1alpha1.TektonConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TektonConfig), err
}

// Update takes the representation of a tektonConfig and updates it. Returns the server's representation of the tektonConfig, and an error, if there is any.
func (c *FakeTektonConfigs) Update(ctx context.Context, tektonConfig *v1alpha1.TektonConfig, opts v1.UpdateOptions) (result *v1alpha1.TektonConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(tektonconfigsResource, tektonConfig), &v1alpha1.TektonConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TektonConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTektonConfigs) UpdateStatus(ctx context.Context, tektonConfig *v1alpha1.TektonConfig, opts v1.UpdateOptions) (*v1alpha1.TektonConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(tektonconfigsResource, "status", tektonConfig), &v1alpha1.TektonConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TektonConfig), err
}

// Delete takes name of the tektonConfig and deletes it. Returns an error if one occurs.
func (c *FakeTektonConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(tektonconfigsResource, name, opts), &v1alpha1.TektonConfig{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTektonConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(tektonconfigsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TektonConfigList{})
	return err
}

// Patch applies the patch and returns the patched tektonConfig.
func (c *FakeTektonConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TektonConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(tektonconfigsResource, name, pt, data, subresources...), &v1alpha1.TektonConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TektonConfig), err
}
