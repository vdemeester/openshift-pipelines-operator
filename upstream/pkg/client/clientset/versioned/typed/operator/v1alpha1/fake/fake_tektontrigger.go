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

// FakeTektonTriggers implements TektonTriggerInterface
type FakeTektonTriggers struct {
	Fake *FakeOperatorV1alpha1
}

var tektontriggersResource = v1alpha1.SchemeGroupVersion.WithResource("tektontriggers")

var tektontriggersKind = v1alpha1.SchemeGroupVersion.WithKind("TektonTrigger")

// Get takes name of the tektonTrigger, and returns the corresponding tektonTrigger object, and an error if there is any.
func (c *FakeTektonTriggers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TektonTrigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(tektontriggersResource, name), &v1alpha1.TektonTrigger{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TektonTrigger), err
}

// List takes label and field selectors, and returns the list of TektonTriggers that match those selectors.
func (c *FakeTektonTriggers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TektonTriggerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(tektontriggersResource, tektontriggersKind, opts), &v1alpha1.TektonTriggerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TektonTriggerList{ListMeta: obj.(*v1alpha1.TektonTriggerList).ListMeta}
	for _, item := range obj.(*v1alpha1.TektonTriggerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tektonTriggers.
func (c *FakeTektonTriggers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(tektontriggersResource, opts))
}

// Create takes the representation of a tektonTrigger and creates it.  Returns the server's representation of the tektonTrigger, and an error, if there is any.
func (c *FakeTektonTriggers) Create(ctx context.Context, tektonTrigger *v1alpha1.TektonTrigger, opts v1.CreateOptions) (result *v1alpha1.TektonTrigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(tektontriggersResource, tektonTrigger), &v1alpha1.TektonTrigger{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TektonTrigger), err
}

// Update takes the representation of a tektonTrigger and updates it. Returns the server's representation of the tektonTrigger, and an error, if there is any.
func (c *FakeTektonTriggers) Update(ctx context.Context, tektonTrigger *v1alpha1.TektonTrigger, opts v1.UpdateOptions) (result *v1alpha1.TektonTrigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(tektontriggersResource, tektonTrigger), &v1alpha1.TektonTrigger{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TektonTrigger), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTektonTriggers) UpdateStatus(ctx context.Context, tektonTrigger *v1alpha1.TektonTrigger, opts v1.UpdateOptions) (*v1alpha1.TektonTrigger, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(tektontriggersResource, "status", tektonTrigger), &v1alpha1.TektonTrigger{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TektonTrigger), err
}

// Delete takes name of the tektonTrigger and deletes it. Returns an error if one occurs.
func (c *FakeTektonTriggers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(tektontriggersResource, name, opts), &v1alpha1.TektonTrigger{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTektonTriggers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(tektontriggersResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TektonTriggerList{})
	return err
}

// Patch applies the patch and returns the patched tektonTrigger.
func (c *FakeTektonTriggers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TektonTrigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(tektontriggersResource, name, pt, data, subresources...), &v1alpha1.TektonTrigger{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TektonTrigger), err
}
