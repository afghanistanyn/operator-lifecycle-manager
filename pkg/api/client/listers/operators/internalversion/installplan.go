/*
Copyright 2020 Red Hat, Inc.

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

package internalversion

import (
	operators "github.com/operator-framework/operator-lifecycle-manager/pkg/api/apis/operators"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// InstallPlanLister helps list InstallPlans.
type InstallPlanLister interface {
	// List lists all InstallPlans in the indexer.
	List(selector labels.Selector) (ret []*operators.InstallPlan, err error)
	// InstallPlans returns an object that can list and get InstallPlans.
	InstallPlans(namespace string) InstallPlanNamespaceLister
	InstallPlanListerExpansion
}

// installPlanLister implements the InstallPlanLister interface.
type installPlanLister struct {
	indexer cache.Indexer
}

// NewInstallPlanLister returns a new InstallPlanLister.
func NewInstallPlanLister(indexer cache.Indexer) InstallPlanLister {
	return &installPlanLister{indexer: indexer}
}

// List lists all InstallPlans in the indexer.
func (s *installPlanLister) List(selector labels.Selector) (ret []*operators.InstallPlan, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*operators.InstallPlan))
	})
	return ret, err
}

// InstallPlans returns an object that can list and get InstallPlans.
func (s *installPlanLister) InstallPlans(namespace string) InstallPlanNamespaceLister {
	return installPlanNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InstallPlanNamespaceLister helps list and get InstallPlans.
type InstallPlanNamespaceLister interface {
	// List lists all InstallPlans in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*operators.InstallPlan, err error)
	// Get retrieves the InstallPlan from the indexer for a given namespace and name.
	Get(name string) (*operators.InstallPlan, error)
	InstallPlanNamespaceListerExpansion
}

// installPlanNamespaceLister implements the InstallPlanNamespaceLister
// interface.
type installPlanNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all InstallPlans in the indexer for a given namespace.
func (s installPlanNamespaceLister) List(selector labels.Selector) (ret []*operators.InstallPlan, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*operators.InstallPlan))
	})
	return ret, err
}

// Get retrieves the InstallPlan from the indexer for a given namespace and name.
func (s installPlanNamespaceLister) Get(name string) (*operators.InstallPlan, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(operators.Resource("installplan"), name)
	}
	return obj.(*operators.InstallPlan), nil
}
