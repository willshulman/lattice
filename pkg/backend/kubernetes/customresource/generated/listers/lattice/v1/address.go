// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/mlab-lattice/lattice/pkg/backend/kubernetes/customresource/apis/lattice/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AddressLister helps list Addresses.
type AddressLister interface {
	// List lists all Addresses in the indexer.
	List(selector labels.Selector) (ret []*v1.Address, err error)
	// Addresses returns an object that can list and get Addresses.
	Addresses(namespace string) AddressNamespaceLister
	AddressListerExpansion
}

// addressLister implements the AddressLister interface.
type addressLister struct {
	indexer cache.Indexer
}

// NewAddressLister returns a new AddressLister.
func NewAddressLister(indexer cache.Indexer) AddressLister {
	return &addressLister{indexer: indexer}
}

// List lists all Addresses in the indexer.
func (s *addressLister) List(selector labels.Selector) (ret []*v1.Address, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Address))
	})
	return ret, err
}

// Addresses returns an object that can list and get Addresses.
func (s *addressLister) Addresses(namespace string) AddressNamespaceLister {
	return addressNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AddressNamespaceLister helps list and get Addresses.
type AddressNamespaceLister interface {
	// List lists all Addresses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.Address, err error)
	// Get retrieves the Address from the indexer for a given namespace and name.
	Get(name string) (*v1.Address, error)
	AddressNamespaceListerExpansion
}

// addressNamespaceLister implements the AddressNamespaceLister
// interface.
type addressNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Addresses in the indexer for a given namespace.
func (s addressNamespaceLister) List(selector labels.Selector) (ret []*v1.Address, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Address))
	})
	return ret, err
}

// Get retrieves the Address from the indexer for a given namespace and name.
func (s addressNamespaceLister) Get(name string) (*v1.Address, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("address"), name)
	}
	return obj.(*v1.Address), nil
}
