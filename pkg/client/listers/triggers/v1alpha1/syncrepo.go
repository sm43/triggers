/*
Copyright 2019 The Tekton Authors

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
	v1alpha1 "github.com/tektoncd/triggers/pkg/apis/triggers/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SyncRepoLister helps list SyncRepos.
// All objects returned here must be treated as read-only.
type SyncRepoLister interface {
	// List lists all SyncRepos in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SyncRepo, err error)
	// SyncRepos returns an object that can list and get SyncRepos.
	SyncRepos(namespace string) SyncRepoNamespaceLister
	SyncRepoListerExpansion
}

// syncRepoLister implements the SyncRepoLister interface.
type syncRepoLister struct {
	indexer cache.Indexer
}

// NewSyncRepoLister returns a new SyncRepoLister.
func NewSyncRepoLister(indexer cache.Indexer) SyncRepoLister {
	return &syncRepoLister{indexer: indexer}
}

// List lists all SyncRepos in the indexer.
func (s *syncRepoLister) List(selector labels.Selector) (ret []*v1alpha1.SyncRepo, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SyncRepo))
	})
	return ret, err
}

// SyncRepos returns an object that can list and get SyncRepos.
func (s *syncRepoLister) SyncRepos(namespace string) SyncRepoNamespaceLister {
	return syncRepoNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SyncRepoNamespaceLister helps list and get SyncRepos.
// All objects returned here must be treated as read-only.
type SyncRepoNamespaceLister interface {
	// List lists all SyncRepos in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SyncRepo, err error)
	// Get retrieves the SyncRepo from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SyncRepo, error)
	SyncRepoNamespaceListerExpansion
}

// syncRepoNamespaceLister implements the SyncRepoNamespaceLister
// interface.
type syncRepoNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SyncRepos in the indexer for a given namespace.
func (s syncRepoNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SyncRepo, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SyncRepo))
	})
	return ret, err
}

// Get retrieves the SyncRepo from the indexer for a given namespace and name.
func (s syncRepoNamespaceLister) Get(name string) (*v1alpha1.SyncRepo, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("syncrepo"), name)
	}
	return obj.(*v1alpha1.SyncRepo), nil
}
