// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/mlab-lattice/lattice/pkg/backend/kubernetes/customresource/apis/lattice/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ComponentBuildLister helps list ComponentBuilds.
type ComponentBuildLister interface {
	// List lists all ComponentBuilds in the indexer.
	List(selector labels.Selector) (ret []*v1.ComponentBuild, err error)
	// ComponentBuilds returns an object that can list and get ComponentBuilds.
	ComponentBuilds(namespace string) ComponentBuildNamespaceLister
	ComponentBuildListerExpansion
}

// componentBuildLister implements the ComponentBuildLister interface.
type componentBuildLister struct {
	indexer cache.Indexer
}

// NewComponentBuildLister returns a new ComponentBuildLister.
func NewComponentBuildLister(indexer cache.Indexer) ComponentBuildLister {
	return &componentBuildLister{indexer: indexer}
}

// List lists all ComponentBuilds in the indexer.
func (s *componentBuildLister) List(selector labels.Selector) (ret []*v1.ComponentBuild, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ComponentBuild))
	})
	return ret, err
}

// ComponentBuilds returns an object that can list and get ComponentBuilds.
func (s *componentBuildLister) ComponentBuilds(namespace string) ComponentBuildNamespaceLister {
	return componentBuildNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ComponentBuildNamespaceLister helps list and get ComponentBuilds.
type ComponentBuildNamespaceLister interface {
	// List lists all ComponentBuilds in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.ComponentBuild, err error)
	// Get retrieves the ComponentBuild from the indexer for a given namespace and name.
	Get(name string) (*v1.ComponentBuild, error)
	ComponentBuildNamespaceListerExpansion
}

// componentBuildNamespaceLister implements the ComponentBuildNamespaceLister
// interface.
type componentBuildNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ComponentBuilds in the indexer for a given namespace.
func (s componentBuildNamespaceLister) List(selector labels.Selector) (ret []*v1.ComponentBuild, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ComponentBuild))
	})
	return ret, err
}

// Get retrieves the ComponentBuild from the indexer for a given namespace and name.
func (s componentBuildNamespaceLister) Get(name string) (*v1.ComponentBuild, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("componentbuild"), name)
	}
	return obj.(*v1.ComponentBuild), nil
}
