// Copyright 2021 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	trainingv1 "github.com/kubeflow/training-operator/pkg/apis/training/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePyTorchJobs implements PyTorchJobInterface
type FakePyTorchJobs struct {
	Fake *FakeKubeflowV1
	ns   string
}

var pytorchjobsResource = schema.GroupVersionResource{Group: "kubeflow.org", Version: "v1", Resource: "pytorchjobs"}

var pytorchjobsKind = schema.GroupVersionKind{Group: "kubeflow.org", Version: "v1", Kind: "PyTorchJob"}

// Get takes name of the pyTorchJob, and returns the corresponding pyTorchJob object, and an error if there is any.
func (c *FakePyTorchJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *trainingv1.PyTorchJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pytorchjobsResource, c.ns, name), &trainingv1.PyTorchJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*trainingv1.PyTorchJob), err
}

// List takes label and field selectors, and returns the list of PyTorchJobs that match those selectors.
func (c *FakePyTorchJobs) List(ctx context.Context, opts v1.ListOptions) (result *trainingv1.PyTorchJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pytorchjobsResource, pytorchjobsKind, c.ns, opts), &trainingv1.PyTorchJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &trainingv1.PyTorchJobList{ListMeta: obj.(*trainingv1.PyTorchJobList).ListMeta}
	for _, item := range obj.(*trainingv1.PyTorchJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pyTorchJobs.
func (c *FakePyTorchJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pytorchjobsResource, c.ns, opts))

}

// Create takes the representation of a pyTorchJob and creates it.  Returns the server's representation of the pyTorchJob, and an error, if there is any.
func (c *FakePyTorchJobs) Create(ctx context.Context, pyTorchJob *trainingv1.PyTorchJob, opts v1.CreateOptions) (result *trainingv1.PyTorchJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pytorchjobsResource, c.ns, pyTorchJob), &trainingv1.PyTorchJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*trainingv1.PyTorchJob), err
}

// Update takes the representation of a pyTorchJob and updates it. Returns the server's representation of the pyTorchJob, and an error, if there is any.
func (c *FakePyTorchJobs) Update(ctx context.Context, pyTorchJob *trainingv1.PyTorchJob, opts v1.UpdateOptions) (result *trainingv1.PyTorchJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pytorchjobsResource, c.ns, pyTorchJob), &trainingv1.PyTorchJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*trainingv1.PyTorchJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePyTorchJobs) UpdateStatus(ctx context.Context, pyTorchJob *trainingv1.PyTorchJob, opts v1.UpdateOptions) (*trainingv1.PyTorchJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(pytorchjobsResource, "status", c.ns, pyTorchJob), &trainingv1.PyTorchJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*trainingv1.PyTorchJob), err
}

// Delete takes name of the pyTorchJob and deletes it. Returns an error if one occurs.
func (c *FakePyTorchJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(pytorchjobsResource, c.ns, name, opts), &trainingv1.PyTorchJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePyTorchJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pytorchjobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &trainingv1.PyTorchJobList{})
	return err
}

// Patch applies the patch and returns the patched pyTorchJob.
func (c *FakePyTorchJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *trainingv1.PyTorchJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pytorchjobsResource, c.ns, name, pt, data, subresources...), &trainingv1.PyTorchJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*trainingv1.PyTorchJob), err
}
