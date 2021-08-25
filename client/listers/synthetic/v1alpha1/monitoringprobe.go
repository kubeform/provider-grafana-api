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

// MonitoringProbeLister helps list MonitoringProbes.
// All objects returned here must be treated as read-only.
type MonitoringProbeLister interface {
	// List lists all MonitoringProbes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MonitoringProbe, err error)
	// MonitoringProbes returns an object that can list and get MonitoringProbes.
	MonitoringProbes(namespace string) MonitoringProbeNamespaceLister
	MonitoringProbeListerExpansion
}

// monitoringProbeLister implements the MonitoringProbeLister interface.
type monitoringProbeLister struct {
	indexer cache.Indexer
}

// NewMonitoringProbeLister returns a new MonitoringProbeLister.
func NewMonitoringProbeLister(indexer cache.Indexer) MonitoringProbeLister {
	return &monitoringProbeLister{indexer: indexer}
}

// List lists all MonitoringProbes in the indexer.
func (s *monitoringProbeLister) List(selector labels.Selector) (ret []*v1alpha1.MonitoringProbe, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MonitoringProbe))
	})
	return ret, err
}

// MonitoringProbes returns an object that can list and get MonitoringProbes.
func (s *monitoringProbeLister) MonitoringProbes(namespace string) MonitoringProbeNamespaceLister {
	return monitoringProbeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MonitoringProbeNamespaceLister helps list and get MonitoringProbes.
// All objects returned here must be treated as read-only.
type MonitoringProbeNamespaceLister interface {
	// List lists all MonitoringProbes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MonitoringProbe, err error)
	// Get retrieves the MonitoringProbe from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.MonitoringProbe, error)
	MonitoringProbeNamespaceListerExpansion
}

// monitoringProbeNamespaceLister implements the MonitoringProbeNamespaceLister
// interface.
type monitoringProbeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MonitoringProbes in the indexer for a given namespace.
func (s monitoringProbeNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MonitoringProbe, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MonitoringProbe))
	})
	return ret, err
}

// Get retrieves the MonitoringProbe from the indexer for a given namespace and name.
func (s monitoringProbeNamespaceLister) Get(name string) (*v1alpha1.MonitoringProbe, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("monitoringprobe"), name)
	}
	return obj.(*v1alpha1.MonitoringProbe), nil
}