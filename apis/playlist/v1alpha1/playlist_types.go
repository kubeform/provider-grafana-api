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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Playlist struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PlaylistSpec   `json:"spec,omitempty"`
	Status            PlaylistStatus `json:"status,omitempty"`
}

type PlaylistSpecItem struct {
	// +optional
	ID    *string `json:"ID,omitempty" tf:"id"`
	Order *int64  `json:"order" tf:"order"`
	Title *string `json:"title" tf:"title"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type PlaylistSpec struct {
	State *PlaylistSpecResource `json:"state,omitempty" tf:"-"`

	Resource PlaylistSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type PlaylistSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Interval *string            `json:"interval" tf:"interval"`
	Item     []PlaylistSpecItem `json:"item" tf:"item"`
	// The name of the playlist.
	Name *string `json:"name" tf:"name"`
	// +optional
	OrgID *string `json:"orgID,omitempty" tf:"org_id"`
}

type PlaylistStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PlaylistList is a list of Playlists
type PlaylistList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Playlist CRD objects
	Items []Playlist `json:"items,omitempty"`
}