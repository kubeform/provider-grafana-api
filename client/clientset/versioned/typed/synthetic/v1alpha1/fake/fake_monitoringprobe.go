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

	v1alpha1 "kubeform.dev/provider-grafana-api/apis/synthetic/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMonitoringProbes implements MonitoringProbeInterface
type FakeMonitoringProbes struct {
	Fake *FakeSyntheticV1alpha1
	ns   string
}

var monitoringprobesResource = schema.GroupVersionResource{Group: "synthetic.grafana.kubeform.com", Version: "v1alpha1", Resource: "monitoringprobes"}

var monitoringprobesKind = schema.GroupVersionKind{Group: "synthetic.grafana.kubeform.com", Version: "v1alpha1", Kind: "MonitoringProbe"}

// Get takes name of the monitoringProbe, and returns the corresponding monitoringProbe object, and an error if there is any.
func (c *FakeMonitoringProbes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MonitoringProbe, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(monitoringprobesResource, c.ns, name), &v1alpha1.MonitoringProbe{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitoringProbe), err
}

// List takes label and field selectors, and returns the list of MonitoringProbes that match those selectors.
func (c *FakeMonitoringProbes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MonitoringProbeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(monitoringprobesResource, monitoringprobesKind, c.ns, opts), &v1alpha1.MonitoringProbeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MonitoringProbeList{ListMeta: obj.(*v1alpha1.MonitoringProbeList).ListMeta}
	for _, item := range obj.(*v1alpha1.MonitoringProbeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested monitoringProbes.
func (c *FakeMonitoringProbes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(monitoringprobesResource, c.ns, opts))

}

// Create takes the representation of a monitoringProbe and creates it.  Returns the server's representation of the monitoringProbe, and an error, if there is any.
func (c *FakeMonitoringProbes) Create(ctx context.Context, monitoringProbe *v1alpha1.MonitoringProbe, opts v1.CreateOptions) (result *v1alpha1.MonitoringProbe, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(monitoringprobesResource, c.ns, monitoringProbe), &v1alpha1.MonitoringProbe{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitoringProbe), err
}

// Update takes the representation of a monitoringProbe and updates it. Returns the server's representation of the monitoringProbe, and an error, if there is any.
func (c *FakeMonitoringProbes) Update(ctx context.Context, monitoringProbe *v1alpha1.MonitoringProbe, opts v1.UpdateOptions) (result *v1alpha1.MonitoringProbe, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(monitoringprobesResource, c.ns, monitoringProbe), &v1alpha1.MonitoringProbe{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitoringProbe), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMonitoringProbes) UpdateStatus(ctx context.Context, monitoringProbe *v1alpha1.MonitoringProbe, opts v1.UpdateOptions) (*v1alpha1.MonitoringProbe, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(monitoringprobesResource, "status", c.ns, monitoringProbe), &v1alpha1.MonitoringProbe{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitoringProbe), err
}

// Delete takes name of the monitoringProbe and deletes it. Returns an error if one occurs.
func (c *FakeMonitoringProbes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(monitoringprobesResource, c.ns, name), &v1alpha1.MonitoringProbe{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMonitoringProbes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(monitoringprobesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MonitoringProbeList{})
	return err
}

// Patch applies the patch and returns the patched monitoringProbe.
func (c *FakeMonitoringProbes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MonitoringProbe, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(monitoringprobesResource, c.ns, name, pt, data, subresources...), &v1alpha1.MonitoringProbe{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitoringProbe), err
}