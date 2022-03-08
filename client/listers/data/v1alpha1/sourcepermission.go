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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/provider-grafana-api/apis/data/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SourcePermissionLister helps list SourcePermissions.
// All objects returned here must be treated as read-only.
type SourcePermissionLister interface {
	// List lists all SourcePermissions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SourcePermission, err error)
	// SourcePermissions returns an object that can list and get SourcePermissions.
	SourcePermissions(namespace string) SourcePermissionNamespaceLister
	SourcePermissionListerExpansion
}

// sourcePermissionLister implements the SourcePermissionLister interface.
type sourcePermissionLister struct {
	indexer cache.Indexer
}

// NewSourcePermissionLister returns a new SourcePermissionLister.
func NewSourcePermissionLister(indexer cache.Indexer) SourcePermissionLister {
	return &sourcePermissionLister{indexer: indexer}
}

// List lists all SourcePermissions in the indexer.
func (s *sourcePermissionLister) List(selector labels.Selector) (ret []*v1alpha1.SourcePermission, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SourcePermission))
	})
	return ret, err
}

// SourcePermissions returns an object that can list and get SourcePermissions.
func (s *sourcePermissionLister) SourcePermissions(namespace string) SourcePermissionNamespaceLister {
	return sourcePermissionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SourcePermissionNamespaceLister helps list and get SourcePermissions.
// All objects returned here must be treated as read-only.
type SourcePermissionNamespaceLister interface {
	// List lists all SourcePermissions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SourcePermission, err error)
	// Get retrieves the SourcePermission from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SourcePermission, error)
	SourcePermissionNamespaceListerExpansion
}

// sourcePermissionNamespaceLister implements the SourcePermissionNamespaceLister
// interface.
type sourcePermissionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SourcePermissions in the indexer for a given namespace.
func (s sourcePermissionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SourcePermission, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SourcePermission))
	})
	return ret, err
}

// Get retrieves the SourcePermission from the indexer for a given namespace and name.
func (s sourcePermissionNamespaceLister) Get(name string) (*v1alpha1.SourcePermission, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sourcepermission"), name)
	}
	return obj.(*v1alpha1.SourcePermission), nil
}