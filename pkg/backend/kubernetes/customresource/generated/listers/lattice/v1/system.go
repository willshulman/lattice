// This file was automatically generated by lister-gen

package v1

import (
	v1 "github.com/mlab-lattice/lattice/pkg/backend/kubernetes/customresource/apis/lattice/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SystemLister helps list Systems.
type SystemLister interface {
	// List lists all Systems in the indexer.
	List(selector labels.Selector) (ret []*v1.System, err error)
	// Systems returns an object that can list and get Systems.
	Systems(namespace string) SystemNamespaceLister
	SystemListerExpansion
}

// systemLister implements the SystemLister interface.
type systemLister struct {
	indexer cache.Indexer
}

// NewSystemLister returns a new SystemLister.
func NewSystemLister(indexer cache.Indexer) SystemLister {
	return &systemLister{indexer: indexer}
}

// List lists all Systems in the indexer.
func (s *systemLister) List(selector labels.Selector) (ret []*v1.System, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.System))
	})
	return ret, err
}

// Systems returns an object that can list and get Systems.
func (s *systemLister) Systems(namespace string) SystemNamespaceLister {
	return systemNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SystemNamespaceLister helps list and get Systems.
type SystemNamespaceLister interface {
	// List lists all Systems in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.System, err error)
	// Get retrieves the System from the indexer for a given namespace and name.
	Get(name string) (*v1.System, error)
	SystemNamespaceListerExpansion
}

// systemNamespaceLister implements the SystemNamespaceLister
// interface.
type systemNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Systems in the indexer for a given namespace.
func (s systemNamespaceLister) List(selector labels.Selector) (ret []*v1.System, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.System))
	})
	return ret, err
}

// Get retrieves the System from the indexer for a given namespace and name.
func (s systemNamespaceLister) Get(name string) (*v1.System, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("system"), name)
	}
	return obj.(*v1.System), nil
}
