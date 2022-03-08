/*
Copyright AppsCode Inc. and Contributors

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

	v1alpha1 "kubeform.dev/provider-grafana-api/apis/machine/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLearningJobs implements LearningJobInterface
type FakeLearningJobs struct {
	Fake *FakeMachineV1alpha1
	ns   string
}

var learningjobsResource = schema.GroupVersionResource{Group: "machine.grafana.kubeform.com", Version: "v1alpha1", Resource: "learningjobs"}

var learningjobsKind = schema.GroupVersionKind{Group: "machine.grafana.kubeform.com", Version: "v1alpha1", Kind: "LearningJob"}

// Get takes name of the learningJob, and returns the corresponding learningJob object, and an error if there is any.
func (c *FakeLearningJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LearningJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(learningjobsResource, c.ns, name), &v1alpha1.LearningJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LearningJob), err
}

// List takes label and field selectors, and returns the list of LearningJobs that match those selectors.
func (c *FakeLearningJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LearningJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(learningjobsResource, learningjobsKind, c.ns, opts), &v1alpha1.LearningJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LearningJobList{ListMeta: obj.(*v1alpha1.LearningJobList).ListMeta}
	for _, item := range obj.(*v1alpha1.LearningJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested learningJobs.
func (c *FakeLearningJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(learningjobsResource, c.ns, opts))

}

// Create takes the representation of a learningJob and creates it.  Returns the server's representation of the learningJob, and an error, if there is any.
func (c *FakeLearningJobs) Create(ctx context.Context, learningJob *v1alpha1.LearningJob, opts v1.CreateOptions) (result *v1alpha1.LearningJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(learningjobsResource, c.ns, learningJob), &v1alpha1.LearningJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LearningJob), err
}

// Update takes the representation of a learningJob and updates it. Returns the server's representation of the learningJob, and an error, if there is any.
func (c *FakeLearningJobs) Update(ctx context.Context, learningJob *v1alpha1.LearningJob, opts v1.UpdateOptions) (result *v1alpha1.LearningJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(learningjobsResource, c.ns, learningJob), &v1alpha1.LearningJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LearningJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLearningJobs) UpdateStatus(ctx context.Context, learningJob *v1alpha1.LearningJob, opts v1.UpdateOptions) (*v1alpha1.LearningJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(learningjobsResource, "status", c.ns, learningJob), &v1alpha1.LearningJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LearningJob), err
}

// Delete takes name of the learningJob and deletes it. Returns an error if one occurs.
func (c *FakeLearningJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(learningjobsResource, c.ns, name), &v1alpha1.LearningJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLearningJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(learningjobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LearningJobList{})
	return err
}

// Patch applies the patch and returns the patched learningJob.
func (c *FakeLearningJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LearningJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(learningjobsResource, c.ns, name, pt, data, subresources...), &v1alpha1.LearningJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LearningJob), err
}
