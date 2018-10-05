// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/mlab-lattice/lattice/pkg/backend/kubernetes/customresource/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Addresses returns a AddressInformer.
	Addresses() AddressInformer
	// Builds returns a BuildInformer.
	Builds() BuildInformer
	// Configs returns a ConfigInformer.
	Configs() ConfigInformer
	// ContainerBuilds returns a ContainerBuildInformer.
	ContainerBuilds() ContainerBuildInformer
	// Deploys returns a DeployInformer.
	Deploys() DeployInformer
	// GitTemplates returns a GitTemplateInformer.
	GitTemplates() GitTemplateInformer
	// Jobs returns a JobInformer.
	Jobs() JobInformer
	// NodePools returns a NodePoolInformer.
	NodePools() NodePoolInformer
	// Services returns a ServiceInformer.
	Services() ServiceInformer
	// Systems returns a SystemInformer.
	Systems() SystemInformer
	// Teardowns returns a TeardownInformer.
	Teardowns() TeardownInformer
	// Templates returns a TemplateInformer.
	Templates() TemplateInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Addresses returns a AddressInformer.
func (v *version) Addresses() AddressInformer {
	return &addressInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Builds returns a BuildInformer.
func (v *version) Builds() BuildInformer {
	return &buildInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Configs returns a ConfigInformer.
func (v *version) Configs() ConfigInformer {
	return &configInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ContainerBuilds returns a ContainerBuildInformer.
func (v *version) ContainerBuilds() ContainerBuildInformer {
	return &containerBuildInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Deploys returns a DeployInformer.
func (v *version) Deploys() DeployInformer {
	return &deployInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// GitTemplates returns a GitTemplateInformer.
func (v *version) GitTemplates() GitTemplateInformer {
	return &gitTemplateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Jobs returns a JobInformer.
func (v *version) Jobs() JobInformer {
	return &jobInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NodePools returns a NodePoolInformer.
func (v *version) NodePools() NodePoolInformer {
	return &nodePoolInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Services returns a ServiceInformer.
func (v *version) Services() ServiceInformer {
	return &serviceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Systems returns a SystemInformer.
func (v *version) Systems() SystemInformer {
	return &systemInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Teardowns returns a TeardownInformer.
func (v *version) Teardowns() TeardownInformer {
	return &teardownInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Templates returns a TemplateInformer.
func (v *version) Templates() TemplateInformer {
	return &templateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
