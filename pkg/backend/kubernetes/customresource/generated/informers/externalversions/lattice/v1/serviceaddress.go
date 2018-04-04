// This file was automatically generated by informer-gen

package v1

import (
	lattice_v1 "github.com/mlab-lattice/lattice/pkg/backend/kubernetes/customresource/apis/lattice/v1"
	versioned "github.com/mlab-lattice/lattice/pkg/backend/kubernetes/customresource/generated/clientset/versioned"
	internalinterfaces "github.com/mlab-lattice/lattice/pkg/backend/kubernetes/customresource/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/mlab-lattice/lattice/pkg/backend/kubernetes/customresource/generated/listers/lattice/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// ServiceAddressInformer provides access to a shared informer and lister for
// ServiceAddresses.
type ServiceAddressInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ServiceAddressLister
}

type serviceAddressInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewServiceAddressInformer constructs a new informer for ServiceAddress type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServiceAddressInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredServiceAddressInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredServiceAddressInformer constructs a new informer for ServiceAddress type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServiceAddressInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LatticeV1().ServiceAddresses(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LatticeV1().ServiceAddresses(namespace).Watch(options)
			},
		},
		&lattice_v1.ServiceAddress{},
		resyncPeriod,
		indexers,
	)
}

func (f *serviceAddressInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredServiceAddressInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *serviceAddressInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&lattice_v1.ServiceAddress{}, f.defaultInformer)
}

func (f *serviceAddressInformer) Lister() v1.ServiceAddressLister {
	return v1.NewServiceAddressLister(f.Informer().GetIndexer())
}
