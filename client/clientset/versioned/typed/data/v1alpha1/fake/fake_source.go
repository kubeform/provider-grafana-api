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

	v1alpha1 "kubeform.dev/provider-grafana-api/apis/data/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSources implements SourceInterface
type FakeSources struct {
	Fake *FakeDataV1alpha1
	ns   string
}

var sourcesResource = schema.GroupVersionResource{Group: "data.grafana.kubeform.com", Version: "v1alpha1", Resource: "sources"}

var sourcesKind = schema.GroupVersionKind{Group: "data.grafana.kubeform.com", Version: "v1alpha1", Kind: "Source"}

// Get takes name of the source, and returns the corresponding source object, and an error if there is any.
func (c *FakeSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Source, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sourcesResource, c.ns, name), &v1alpha1.Source{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Source), err
}

// List takes label and field selectors, and returns the list of Sources that match those selectors.
func (c *FakeSources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sourcesResource, sourcesKind, c.ns, opts), &v1alpha1.SourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SourceList{ListMeta: obj.(*v1alpha1.SourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.SourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sources.
func (c *FakeSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sourcesResource, c.ns, opts))

}

// Create takes the representation of a source and creates it.  Returns the server's representation of the source, and an error, if there is any.
func (c *FakeSources) Create(ctx context.Context, source *v1alpha1.Source, opts v1.CreateOptions) (result *v1alpha1.Source, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sourcesResource, c.ns, source), &v1alpha1.Source{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Source), err
}

// Update takes the representation of a source and updates it. Returns the server's representation of the source, and an error, if there is any.
func (c *FakeSources) Update(ctx context.Context, source *v1alpha1.Source, opts v1.UpdateOptions) (result *v1alpha1.Source, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sourcesResource, c.ns, source), &v1alpha1.Source{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Source), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSources) UpdateStatus(ctx context.Context, source *v1alpha1.Source, opts v1.UpdateOptions) (*v1alpha1.Source, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sourcesResource, "status", c.ns, source), &v1alpha1.Source{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Source), err
}

// Delete takes name of the source and deletes it. Returns an error if one occurs.
func (c *FakeSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sourcesResource, c.ns, name), &v1alpha1.Source{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sourcesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SourceList{})
	return err
}

// Patch applies the patch and returns the patched source.
func (c *FakeSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Source, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sourcesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Source{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Source), err
}
