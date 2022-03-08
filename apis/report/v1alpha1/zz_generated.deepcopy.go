//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Report) DeepCopyInto(out *Report) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Report.
func (in *Report) DeepCopy() *Report {
	if in == nil {
		return nil
	}
	out := new(Report)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Report) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportList) DeepCopyInto(out *ReportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Report, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportList.
func (in *ReportList) DeepCopy() *ReportList {
	if in == nil {
		return nil
	}
	out := new(ReportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportSpec) DeepCopyInto(out *ReportSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ReportSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportSpec.
func (in *ReportSpec) DeepCopy() *ReportSpec {
	if in == nil {
		return nil
	}
	out := new(ReportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportSpecResource) DeepCopyInto(out *ReportSpecResource) {
	*out = *in
	if in.DashboardID != nil {
		in, out := &in.DashboardID, &out.DashboardID
		*out = new(int64)
		**out = **in
	}
	if in.IncludeDashboardLink != nil {
		in, out := &in.IncludeDashboardLink, &out.IncludeDashboardLink
		*out = new(bool)
		**out = **in
	}
	if in.IncludeTableCsv != nil {
		in, out := &in.IncludeTableCsv, &out.IncludeTableCsv
		*out = new(bool)
		**out = **in
	}
	if in.Layout != nil {
		in, out := &in.Layout, &out.Layout
		*out = new(string)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Orientation != nil {
		in, out := &in.Orientation, &out.Orientation
		*out = new(string)
		**out = **in
	}
	if in.Recipients != nil {
		in, out := &in.Recipients, &out.Recipients
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ReplyTo != nil {
		in, out := &in.ReplyTo, &out.ReplyTo
		*out = new(string)
		**out = **in
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(ReportSpecSchedule)
		(*in).DeepCopyInto(*out)
	}
	if in.TimeRange != nil {
		in, out := &in.TimeRange, &out.TimeRange
		*out = new(ReportSpecTimeRange)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportSpecResource.
func (in *ReportSpecResource) DeepCopy() *ReportSpecResource {
	if in == nil {
		return nil
	}
	out := new(ReportSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportSpecSchedule) DeepCopyInto(out *ReportSpecSchedule) {
	*out = *in
	if in.CustomInterval != nil {
		in, out := &in.CustomInterval, &out.CustomInterval
		*out = new(string)
		**out = **in
	}
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = new(string)
		**out = **in
	}
	if in.Frequency != nil {
		in, out := &in.Frequency, &out.Frequency
		*out = new(string)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(string)
		**out = **in
	}
	if in.WorkdaysOnly != nil {
		in, out := &in.WorkdaysOnly, &out.WorkdaysOnly
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportSpecSchedule.
func (in *ReportSpecSchedule) DeepCopy() *ReportSpecSchedule {
	if in == nil {
		return nil
	}
	out := new(ReportSpecSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportSpecTimeRange) DeepCopyInto(out *ReportSpecTimeRange) {
	*out = *in
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new(string)
		**out = **in
	}
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportSpecTimeRange.
func (in *ReportSpecTimeRange) DeepCopy() *ReportSpecTimeRange {
	if in == nil {
		return nil
	}
	out := new(ReportSpecTimeRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportStatus) DeepCopyInto(out *ReportStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportStatus.
func (in *ReportStatus) DeepCopy() *ReportStatus {
	if in == nil {
		return nil
	}
	out := new(ReportStatus)
	in.DeepCopyInto(out)
	return out
}