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
	v1alpha1 "kubeform.dev/provider-grafana-api/apis/synthetic/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MonitoringCheckLister helps list MonitoringChecks.
// All objects returned here must be treated as read-only.
type MonitoringCheckLister interface {
	// List lists all MonitoringChecks in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MonitoringCheck, err error)
	// MonitoringChecks returns an object that can list and get MonitoringChecks.
	MonitoringChecks(namespace string) MonitoringCheckNamespaceLister
	MonitoringCheckListerExpansion
}

// monitoringCheckLister implements the MonitoringCheckLister interface.
type monitoringCheckLister struct {
	indexer cache.Indexer
}

// NewMonitoringCheckLister returns a new MonitoringCheckLister.
func NewMonitoringCheckLister(indexer cache.Indexer) MonitoringCheckLister {
	return &monitoringCheckLister{indexer: indexer}
}

// List lists all MonitoringChecks in the indexer.
func (s *monitoringCheckLister) List(selector labels.Selector) (ret []*v1alpha1.MonitoringCheck, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MonitoringCheck))
	})
	return ret, err
}

// MonitoringChecks returns an object that can list and get MonitoringChecks.
func (s *monitoringCheckLister) MonitoringChecks(namespace string) MonitoringCheckNamespaceLister {
	return monitoringCheckNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MonitoringCheckNamespaceLister helps list and get MonitoringChecks.
// All objects returned here must be treated as read-only.
type MonitoringCheckNamespaceLister interface {
	// List lists all MonitoringChecks in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MonitoringCheck, err error)
	// Get retrieves the MonitoringCheck from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.MonitoringCheck, error)
	MonitoringCheckNamespaceListerExpansion
}

// monitoringCheckNamespaceLister implements the MonitoringCheckNamespaceLister
// interface.
type monitoringCheckNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MonitoringChecks in the indexer for a given namespace.
func (s monitoringCheckNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MonitoringCheck, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MonitoringCheck))
	})
	return ret, err
}

// Get retrieves the MonitoringCheck from the indexer for a given namespace and name.
func (s monitoringCheckNamespaceLister) Get(name string) (*v1alpha1.MonitoringCheck, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("monitoringcheck"), name)
	}
	return obj.(*v1alpha1.MonitoringCheck), nil
}