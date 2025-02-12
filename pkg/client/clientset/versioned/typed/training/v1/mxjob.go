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

package v1

import (
	"context"
	"time"

	v1 "github.com/kubeflow/training-operator/pkg/apis/training/v1"
	scheme "github.com/kubeflow/training-operator/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MXJobsGetter has a method to return a MXJobInterface.
// A group's client should implement this interface.
type MXJobsGetter interface {
	MXJobs(namespace string) MXJobInterface
}

// MXJobInterface has methods to work with MXJob resources.
type MXJobInterface interface {
	Create(ctx context.Context, mXJob *v1.MXJob, opts metav1.CreateOptions) (*v1.MXJob, error)
	Update(ctx context.Context, mXJob *v1.MXJob, opts metav1.UpdateOptions) (*v1.MXJob, error)
	UpdateStatus(ctx context.Context, mXJob *v1.MXJob, opts metav1.UpdateOptions) (*v1.MXJob, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.MXJob, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.MXJobList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.MXJob, err error)
	MXJobExpansion
}

// mXJobs implements MXJobInterface
type mXJobs struct {
	client rest.Interface
	ns     string
}

// newMXJobs returns a MXJobs
func newMXJobs(c *KubeflowV1Client, namespace string) *mXJobs {
	return &mXJobs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the mXJob, and returns the corresponding mXJob object, and an error if there is any.
func (c *mXJobs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.MXJob, err error) {
	result = &v1.MXJob{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mxjobs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MXJobs that match those selectors.
func (c *mXJobs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.MXJobList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.MXJobList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mxjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested mXJobs.
func (c *mXJobs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mxjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a mXJob and creates it.  Returns the server's representation of the mXJob, and an error, if there is any.
func (c *mXJobs) Create(ctx context.Context, mXJob *v1.MXJob, opts metav1.CreateOptions) (result *v1.MXJob, err error) {
	result = &v1.MXJob{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mxjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(mXJob).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a mXJob and updates it. Returns the server's representation of the mXJob, and an error, if there is any.
func (c *mXJobs) Update(ctx context.Context, mXJob *v1.MXJob, opts metav1.UpdateOptions) (result *v1.MXJob, err error) {
	result = &v1.MXJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mxjobs").
		Name(mXJob.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(mXJob).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *mXJobs) UpdateStatus(ctx context.Context, mXJob *v1.MXJob, opts metav1.UpdateOptions) (result *v1.MXJob, err error) {
	result = &v1.MXJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mxjobs").
		Name(mXJob.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(mXJob).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the mXJob and deletes it. Returns an error if one occurs.
func (c *mXJobs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mxjobs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *mXJobs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mxjobs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched mXJob.
func (c *mXJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.MXJob, err error) {
	result = &v1.MXJob{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mxjobs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
