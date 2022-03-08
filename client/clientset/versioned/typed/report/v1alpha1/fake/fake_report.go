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

	v1alpha1 "kubeform.dev/provider-grafana-api/apis/report/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeReports implements ReportInterface
type FakeReports struct {
	Fake *FakeReportV1alpha1
	ns   string
}

var reportsResource = schema.GroupVersionResource{Group: "report.grafana.kubeform.com", Version: "v1alpha1", Resource: "reports"}

var reportsKind = schema.GroupVersionKind{Group: "report.grafana.kubeform.com", Version: "v1alpha1", Kind: "Report"}

// Get takes name of the report, and returns the corresponding report object, and an error if there is any.
func (c *FakeReports) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Report, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(reportsResource, c.ns, name), &v1alpha1.Report{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Report), err
}

// List takes label and field selectors, and returns the list of Reports that match those selectors.
func (c *FakeReports) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ReportList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(reportsResource, reportsKind, c.ns, opts), &v1alpha1.ReportList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ReportList{ListMeta: obj.(*v1alpha1.ReportList).ListMeta}
	for _, item := range obj.(*v1alpha1.ReportList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested reports.
func (c *FakeReports) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(reportsResource, c.ns, opts))

}

// Create takes the representation of a report and creates it.  Returns the server's representation of the report, and an error, if there is any.
func (c *FakeReports) Create(ctx context.Context, report *v1alpha1.Report, opts v1.CreateOptions) (result *v1alpha1.Report, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(reportsResource, c.ns, report), &v1alpha1.Report{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Report), err
}

// Update takes the representation of a report and updates it. Returns the server's representation of the report, and an error, if there is any.
func (c *FakeReports) Update(ctx context.Context, report *v1alpha1.Report, opts v1.UpdateOptions) (result *v1alpha1.Report, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(reportsResource, c.ns, report), &v1alpha1.Report{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Report), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeReports) UpdateStatus(ctx context.Context, report *v1alpha1.Report, opts v1.UpdateOptions) (*v1alpha1.Report, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(reportsResource, "status", c.ns, report), &v1alpha1.Report{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Report), err
}

// Delete takes name of the report and deletes it. Returns an error if one occurs.
func (c *FakeReports) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(reportsResource, c.ns, name), &v1alpha1.Report{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeReports) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(reportsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ReportList{})
	return err
}

// Patch applies the patch and returns the patched report.
func (c *FakeReports) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Report, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(reportsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Report{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Report), err
}
