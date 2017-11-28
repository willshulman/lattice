package app

import (
	"fmt"

	"github.com/mlab-lattice/system/pkg/kubernetes/constants"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1 "k8s.io/api/core/v1"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"

	"k8s.io/client-go/kubernetes"
)

func seedLatticeSystemEnvironmentManagerAPI(kubeClientset *kubernetes.Clientset) {
	fmt.Println("Seeding lattice-system-environment-manager...")

	latticeSystemEnvironmentManagerAPIDaemonSet := &extensionsv1beta1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      constants.MasterNodeComponentManagerApi,
			Namespace: constants.NamespaceLatticeInternal,
		},
		Spec: extensionsv1beta1.DaemonSetSpec{
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name: constants.MasterNodeComponentManagerApi,
					Labels: map[string]string{
						constants.MasterNodeLabelComponent: constants.MasterNodeComponentManagerApi,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "api",
							Image: latticeContainerRegistry + "/kubernetes-manager-api-rest",
							Args:  []string{"-port", "80"},
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									HostPort:      80,
									ContainerPort: 80,
								},
							},
						},
					},
					HostNetwork:        true,
					DNSPolicy:          corev1.DNSDefault,
					ServiceAccountName: constants.ServiceAccountManagementApi,
					// Can tolerate the master-node taint in the local case when it's not applied harmlessly
					Tolerations: []corev1.Toleration{
						constants.TolerationMasterNode,
					},
				},
			},
		},
	}

	// FIXME: add NodeSelector for cloud providers
	//switch coretypes.Provider(provider) {
	//case coreconstants.ProviderLocal:
	//
	//default:
	//	panic("unsupported provider")
	//}

	pollKubeResourceCreation(func() (interface{}, error) {
		return kubeClientset.
			ExtensionsV1beta1().
			DaemonSets(string(constants.NamespaceLatticeInternal)).
			Create(latticeSystemEnvironmentManagerAPIDaemonSet)
	})
}