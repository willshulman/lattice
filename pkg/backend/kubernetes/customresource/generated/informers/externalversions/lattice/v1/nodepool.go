// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	lattice_v1 "github.com/mlab-lattice/lattice/pkg/backend/kubernetes/customresource/apis/lattice/v1"
	versioned "github.com/mlab-lattice/lattice/pkg/backend/kubernetes/customresource/generated/clientset/versioned"
	internalinterfaces "github.com/mlab-lattice/lattice/pkg/backend/kubernetes/customresource/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/mlab-lattice/lattice/pkg/backend/kubernetes/customresource/generated/listers/lattice/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// NodePoolInformer provides access to a shared informer and lister for
// NodePools.
type NodePoolInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.NodePoolLister
}

type nodePoolInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewNodePoolInformer constructs a new informer for NodePool type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNodePoolInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNodePoolInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredNodePoolInformer constructs a new informer for NodePool type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNodePoolInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LatticeV1().NodePools(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LatticeV1().NodePools(namespace).Watch(options)
			},
		},
		&lattice_v1.NodePool{},
		resyncPeriod,
		indexers,
	)
}

func (f *nodePoolInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNodePoolInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *nodePoolInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&lattice_v1.NodePool{}, f.defaultInformer)
}

func (f *nodePoolInformer) Lister() v1.NodePoolLister {
	return v1.NewNodePoolLister(f.Informer().GetIndexer())
}
