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

	v1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCustomRuns implements CustomRunInterface
type FakeCustomRuns struct {
	Fake *FakeTektonV1beta1
	ns   string
}

var customrunsResource = schema.GroupVersionResource{Group: "tekton.dev", Version: "v1beta1", Resource: "customruns"}

var customrunsKind = schema.GroupVersionKind{Group: "tekton.dev", Version: "v1beta1", Kind: "CustomRun"}

// Get takes name of the customRun, and returns the corresponding customRun object, and an error if there is any.
func (c *FakeCustomRuns) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.CustomRun, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(customrunsResource, c.ns, name), &v1beta1.CustomRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CustomRun), err
}

// List takes label and field selectors, and returns the list of CustomRuns that match those selectors.
func (c *FakeCustomRuns) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.CustomRunList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(customrunsResource, customrunsKind, c.ns, opts), &v1beta1.CustomRunList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.CustomRunList{ListMeta: obj.(*v1beta1.CustomRunList).ListMeta}
	for _, item := range obj.(*v1beta1.CustomRunList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested customRuns.
func (c *FakeCustomRuns) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(customrunsResource, c.ns, opts))

}

// Create takes the representation of a customRun and creates it.  Returns the server's representation of the customRun, and an error, if there is any.
func (c *FakeCustomRuns) Create(ctx context.Context, customRun *v1beta1.CustomRun, opts v1.CreateOptions) (result *v1beta1.CustomRun, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(customrunsResource, c.ns, customRun), &v1beta1.CustomRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CustomRun), err
}

// Update takes the representation of a customRun and updates it. Returns the server's representation of the customRun, and an error, if there is any.
func (c *FakeCustomRuns) Update(ctx context.Context, customRun *v1beta1.CustomRun, opts v1.UpdateOptions) (result *v1beta1.CustomRun, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(customrunsResource, c.ns, customRun), &v1beta1.CustomRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CustomRun), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCustomRuns) UpdateStatus(ctx context.Context, customRun *v1beta1.CustomRun, opts v1.UpdateOptions) (*v1beta1.CustomRun, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(customrunsResource, "status", c.ns, customRun), &v1beta1.CustomRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CustomRun), err
}

// Delete takes name of the customRun and deletes it. Returns an error if one occurs.
func (c *FakeCustomRuns) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(customrunsResource, c.ns, name, opts), &v1beta1.CustomRun{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCustomRuns) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(customrunsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.CustomRunList{})
	return err
}

// Patch applies the patch and returns the patched customRun.
func (c *FakeCustomRuns) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.CustomRun, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(customrunsResource, c.ns, name, pt, data, subresources...), &v1beta1.CustomRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CustomRun), err
}
