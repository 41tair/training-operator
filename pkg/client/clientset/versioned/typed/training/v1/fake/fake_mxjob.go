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

// FakeMXJobs implements MXJobInterface
type FakeMXJobs struct {
	Fake *FakeKubeflowV1
	ns   string
}

var mxjobsResource = schema.GroupVersionResource{Group: "kubeflow.org", Version: "v1", Resource: "mxjobs"}

var mxjobsKind = schema.GroupVersionKind{Group: "kubeflow.org", Version: "v1", Kind: "MXJob"}

// Get takes name of the mXJob, and returns the corresponding mXJob object, and an error if there is any.
func (c *FakeMXJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *trainingv1.MXJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mxjobsResource, c.ns, name), &trainingv1.MXJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*trainingv1.MXJob), err
}

// List takes label and field selectors, and returns the list of MXJobs that match those selectors.
func (c *FakeMXJobs) List(ctx context.Context, opts v1.ListOptions) (result *trainingv1.MXJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mxjobsResource, mxjobsKind, c.ns, opts), &trainingv1.MXJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &trainingv1.MXJobList{ListMeta: obj.(*trainingv1.MXJobList).ListMeta}
	for _, item := range obj.(*trainingv1.MXJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mXJobs.
func (c *FakeMXJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mxjobsResource, c.ns, opts))

}

// Create takes the representation of a mXJob and creates it.  Returns the server's representation of the mXJob, and an error, if there is any.
func (c *FakeMXJobs) Create(ctx context.Context, mXJob *trainingv1.MXJob, opts v1.CreateOptions) (result *trainingv1.MXJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mxjobsResource, c.ns, mXJob), &trainingv1.MXJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*trainingv1.MXJob), err
}

// Update takes the representation of a mXJob and updates it. Returns the server's representation of the mXJob, and an error, if there is any.
func (c *FakeMXJobs) Update(ctx context.Context, mXJob *trainingv1.MXJob, opts v1.UpdateOptions) (result *trainingv1.MXJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mxjobsResource, c.ns, mXJob), &trainingv1.MXJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*trainingv1.MXJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMXJobs) UpdateStatus(ctx context.Context, mXJob *trainingv1.MXJob, opts v1.UpdateOptions) (*trainingv1.MXJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(mxjobsResource, "status", c.ns, mXJob), &trainingv1.MXJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*trainingv1.MXJob), err
}

// Delete takes name of the mXJob and deletes it. Returns an error if one occurs.
func (c *FakeMXJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(mxjobsResource, c.ns, name, opts), &trainingv1.MXJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMXJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mxjobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &trainingv1.MXJobList{})
	return err
}

// Patch applies the patch and returns the patched mXJob.
func (c *FakeMXJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *trainingv1.MXJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mxjobsResource, c.ns, name, pt, data, subresources...), &trainingv1.MXJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*trainingv1.MXJob), err
}
