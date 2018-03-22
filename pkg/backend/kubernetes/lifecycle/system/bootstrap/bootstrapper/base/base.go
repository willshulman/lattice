package base

import (
	kubeconstants "github.com/mlab-lattice/system/pkg/backend/kubernetes/constants"
	latticev1 "github.com/mlab-lattice/system/pkg/backend/kubernetes/customresource/apis/lattice/v1"
	"github.com/mlab-lattice/system/pkg/backend/kubernetes/lifecycle/system/bootstrap/bootstrapper"
	kubeutil "github.com/mlab-lattice/system/pkg/backend/kubernetes/util/kubernetes"
	"github.com/mlab-lattice/system/pkg/definition/tree"
	"github.com/mlab-lattice/system/pkg/types"

	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Options struct {
	LatticeID     types.LatticeID
	SystemID      types.SystemID
	DefinitionURL string
}

func NewBootstrapper(options *Options) *DefaultBootstrapper {
	return &DefaultBootstrapper{
		latticeID:     options.LatticeID,
		systemID:      options.SystemID,
		definitionURL: options.DefinitionURL,
	}
}

type DefaultBootstrapper struct {
	latticeID     types.LatticeID
	systemID      types.SystemID
	definitionURL string
}

func (b *DefaultBootstrapper) BootstrapSystemResources(resources *bootstrapper.SystemResources) {
	namespace := &corev1.Namespace{
		// Include TypeMeta so if this is a dry run it will be printed out
		TypeMeta: metav1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: kubeutil.SystemNamespace(b.latticeID, b.systemID),
			Labels: map[string]string{
				kubeconstants.LabelKeyLatticeID: string(b.latticeID),
			},
		},
	}

	componentBuilderSA := &corev1.ServiceAccount{
		// Include TypeMeta so if this is a dry run it will be printed out
		TypeMeta: metav1.TypeMeta{
			Kind:       "ServiceAccount",
			APIVersion: rbacv1.GroupName + "/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      kubeconstants.ServiceAccountComponentBuilder,
			Namespace: namespace.Name,
		},
	}

	componentBuilderRB := &rbacv1.RoleBinding{
		TypeMeta: metav1.TypeMeta{
			Kind:       "RoleBinding",
			APIVersion: rbacv1.GroupName + "/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      kubeconstants.ControlPlaneServiceComponentBuilder,
			Namespace: componentBuilderSA.Namespace,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      rbacv1.ServiceAccountKind,
				Name:      componentBuilderSA.Name,
				Namespace: componentBuilderSA.Namespace,
			},
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: rbacv1.GroupName,
			Kind:     "ClusterRole",
			Name:     kubeconstants.ControlPlaneServiceComponentBuilder,
		},
	}

	system := &latticev1.System{
		// Include TypeMeta so if this is a dry run it will be printed out
		TypeMeta: metav1.TypeMeta{
			Kind:       "System",
			APIVersion: latticev1.GroupName + "/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      string(b.systemID),
			Namespace: namespace.Name,
			Labels: map[string]string{
				kubeconstants.LabelKeyLatticeID: string(b.latticeID),
			},
		},
		Spec: latticev1.SystemSpec{
			DefinitionURL: b.definitionURL,
			Services:      map[tree.NodePath]latticev1.SystemSpecServiceInfo{},
		},
		Status: latticev1.SystemStatus{
			State: latticev1.SystemStateStable,
		},
	}

	resources.System = system
	resources.Namespace = namespace
	resources.ServiceAccounts = append(resources.ServiceAccounts, componentBuilderSA)
	resources.RoleBindings = append(resources.RoleBindings, componentBuilderRB)
}
