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
	v1alpha1 "kubeform.dev/provider-grafana-api/apis/machine/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LearningJobLister helps list LearningJobs.
// All objects returned here must be treated as read-only.
type LearningJobLister interface {
	// List lists all LearningJobs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LearningJob, err error)
	// LearningJobs returns an object that can list and get LearningJobs.
	LearningJobs(namespace string) LearningJobNamespaceLister
	LearningJobListerExpansion
}

// learningJobLister implements the LearningJobLister interface.
type learningJobLister struct {
	indexer cache.Indexer
}

// NewLearningJobLister returns a new LearningJobLister.
func NewLearningJobLister(indexer cache.Indexer) LearningJobLister {
	return &learningJobLister{indexer: indexer}
}

// List lists all LearningJobs in the indexer.
func (s *learningJobLister) List(selector labels.Selector) (ret []*v1alpha1.LearningJob, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LearningJob))
	})
	return ret, err
}

// LearningJobs returns an object that can list and get LearningJobs.
func (s *learningJobLister) LearningJobs(namespace string) LearningJobNamespaceLister {
	return learningJobNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LearningJobNamespaceLister helps list and get LearningJobs.
// All objects returned here must be treated as read-only.
type LearningJobNamespaceLister interface {
	// List lists all LearningJobs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LearningJob, err error)
	// Get retrieves the LearningJob from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.LearningJob, error)
	LearningJobNamespaceListerExpansion
}

// learningJobNamespaceLister implements the LearningJobNamespaceLister
// interface.
type learningJobNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LearningJobs in the indexer for a given namespace.
func (s learningJobNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LearningJob, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LearningJob))
	})
	return ret, err
}

// Get retrieves the LearningJob from the indexer for a given namespace and name.
func (s learningJobNamespaceLister) Get(name string) (*v1alpha1.LearningJob, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("learningjob"), name)
	}
	return obj.(*v1alpha1.LearningJob), nil
}
