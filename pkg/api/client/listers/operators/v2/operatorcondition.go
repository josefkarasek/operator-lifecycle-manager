/*
Copyright Red Hat, Inc.

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

package v2

import (
	v2 "github.com/operator-framework/api/pkg/operators/v2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// OperatorConditionLister helps list OperatorConditions.
// All objects returned here must be treated as read-only.
type OperatorConditionLister interface {
	// List lists all OperatorConditions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2.OperatorCondition, err error)
	// OperatorConditions returns an object that can list and get OperatorConditions.
	OperatorConditions(namespace string) OperatorConditionNamespaceLister
	OperatorConditionListerExpansion
}

// operatorConditionLister implements the OperatorConditionLister interface.
type operatorConditionLister struct {
	indexer cache.Indexer
}

// NewOperatorConditionLister returns a new OperatorConditionLister.
func NewOperatorConditionLister(indexer cache.Indexer) OperatorConditionLister {
	return &operatorConditionLister{indexer: indexer}
}

// List lists all OperatorConditions in the indexer.
func (s *operatorConditionLister) List(selector labels.Selector) (ret []*v2.OperatorCondition, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.OperatorCondition))
	})
	return ret, err
}

// OperatorConditions returns an object that can list and get OperatorConditions.
func (s *operatorConditionLister) OperatorConditions(namespace string) OperatorConditionNamespaceLister {
	return operatorConditionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// OperatorConditionNamespaceLister helps list and get OperatorConditions.
// All objects returned here must be treated as read-only.
type OperatorConditionNamespaceLister interface {
	// List lists all OperatorConditions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2.OperatorCondition, err error)
	// Get retrieves the OperatorCondition from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v2.OperatorCondition, error)
	OperatorConditionNamespaceListerExpansion
}

// operatorConditionNamespaceLister implements the OperatorConditionNamespaceLister
// interface.
type operatorConditionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all OperatorConditions in the indexer for a given namespace.
func (s operatorConditionNamespaceLister) List(selector labels.Selector) (ret []*v2.OperatorCondition, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.OperatorCondition))
	})
	return ret, err
}

// Get retrieves the OperatorCondition from the indexer for a given namespace and name.
func (s operatorConditionNamespaceLister) Get(name string) (*v2.OperatorCondition, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2.Resource("operatorcondition"), name)
	}
	return obj.(*v2.OperatorCondition), nil
}
