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
	alertv1alpha1 "kubeform.dev/provider-grafana-api/apis/alert/v1alpha1"
	builtinv1alpha1 "kubeform.dev/provider-grafana-api/apis/builtin/v1alpha1"
	dashboardv1alpha1 "kubeform.dev/provider-grafana-api/apis/dashboard/v1alpha1"
	datav1alpha1 "kubeform.dev/provider-grafana-api/apis/data/v1alpha1"
	folderv1alpha1 "kubeform.dev/provider-grafana-api/apis/folder/v1alpha1"
	organizationv1alpha1 "kubeform.dev/provider-grafana-api/apis/organization/v1alpha1"
	rolev1alpha1 "kubeform.dev/provider-grafana-api/apis/role/v1alpha1"
	syntheticv1alpha1 "kubeform.dev/provider-grafana-api/apis/synthetic/v1alpha1"
	teamv1alpha1 "kubeform.dev/provider-grafana-api/apis/team/v1alpha1"
	userv1alpha1 "kubeform.dev/provider-grafana-api/apis/user/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var scheme = runtime.NewScheme()
var codecs = serializer.NewCodecFactory(scheme)

var localSchemeBuilder = runtime.SchemeBuilder{
	alertv1alpha1.AddToScheme,
	builtinv1alpha1.AddToScheme,
	dashboardv1alpha1.AddToScheme,
	datav1alpha1.AddToScheme,
	folderv1alpha1.AddToScheme,
	organizationv1alpha1.AddToScheme,
	rolev1alpha1.AddToScheme,
	syntheticv1alpha1.AddToScheme,
	teamv1alpha1.AddToScheme,
	userv1alpha1.AddToScheme,
}

// AddToScheme adds all types of this clientset into the given scheme. This allows composition
// of clientsets, like in:
//
//   import (
//     "k8s.io/client-go/kubernetes"
//     clientsetscheme "k8s.io/client-go/kubernetes/scheme"
//     aggregatorclientsetscheme "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/scheme"
//   )
//
//   kclientset, _ := kubernetes.NewForConfig(c)
//   _ = aggregatorclientsetscheme.AddToScheme(clientsetscheme.Scheme)
//
// After this, RawExtensions in Kubernetes types will serialize kube-aggregator types
// correctly.
var AddToScheme = localSchemeBuilder.AddToScheme

func init() {
	v1.AddToGroupVersion(scheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(AddToScheme(scheme))
}