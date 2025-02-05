package k8s

import (
	"context"

	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// IngressHelper provides methods for managing Kubernetes Ingresses
type IngressHelper interface {
	Get(ctx context.Context, namespace, name string) (*networkingv1.Ingress, error)
	Create(ctx context.Context, namespace string, ingress *networkingv1.Ingress) (*networkingv1.Ingress, error)
	Update(ctx context.Context, namespace string, ingress *networkingv1.Ingress) (*networkingv1.Ingress, error)
	Delete(ctx context.Context, namespace, name string) error
	List(ctx context.Context, namespace string) (*networkingv1.IngressList, error)
}

type ingressHelper struct {
	client *kubernetes.Clientset
}

// NewIngressHelper creates a new instance of IngressHelper.
//
// Parameters:
//   - client: Kubernetes clientset for interacting with the cluster
//
// Returns:
//   - IngressHelper: Interface for managing Kubernetes Ingresses
func NewIngressHelper(client *kubernetes.Clientset) IngressHelper {
	return &ingressHelper{
		client: client,
	}
}

// Get retrieves an ingress by name from the specified namespace.
//
// Parameters:
//   - ctx: Context for the operation
//   - namespace: Kubernetes namespace
//   - name: Name of the ingress
//
// Returns:
//   - *networkingv1.Ingress: The retrieved ingress
//   - error: nil if successful, error otherwise
func (i *ingressHelper) Get(ctx context.Context, namespace, name string) (*networkingv1.Ingress, error) {
	return i.client.NetworkingV1().Ingresses(namespace).Get(ctx, name, metav1.GetOptions{})
}

// Create creates a new ingress in the specified namespace.
//
// Parameters:
//   - ctx: Context for the operation
//   - namespace: Kubernetes namespace
//   - ingress: Ingress specification to create
//
// Returns:
//   - *networkingv1.Ingress: The created ingress
//   - error: nil if successful, error otherwise
func (i *ingressHelper) Create(ctx context.Context, namespace string, ingress *networkingv1.Ingress) (*networkingv1.Ingress, error) {
	return i.client.NetworkingV1().Ingresses(namespace).Create(ctx, ingress, metav1.CreateOptions{})
}

// Update updates an existing ingress in the specified namespace.
//
// Parameters:
//   - ctx: Context for the operation
//   - namespace: Kubernetes namespace
//   - ingress: Updated ingress specification
//
// Returns:
//   - *networkingv1.Ingress: The updated ingress
//   - error: nil if successful, error otherwise
func (i *ingressHelper) Update(ctx context.Context, namespace string, ingress *networkingv1.Ingress) (*networkingv1.Ingress, error) {
	return i.client.NetworkingV1().Ingresses(namespace).Update(ctx, ingress, metav1.UpdateOptions{})
}

// Delete removes an ingress from the specified namespace.
//
// Parameters:
//   - ctx: Context for the operation
//   - namespace: Kubernetes namespace
//   - name: Name of the ingress to delete
//
// Returns:
//   - error: nil if successful, error otherwise
func (i *ingressHelper) Delete(ctx context.Context, namespace, name string) error {
	return i.client.NetworkingV1().Ingresses(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

// List retrieves all ingresses in the specified namespace.
//
// Parameters:
//   - ctx: Context for the operation
//   - namespace: Kubernetes namespace
//
// Returns:
//   - *networkingv1.IngressList: List of ingresses
//   - error: nil if successful, error otherwise
func (i *ingressHelper) List(ctx context.Context, namespace string) (*networkingv1.IngressList, error) {
	return i.client.NetworkingV1().Ingresses(namespace).List(ctx, metav1.ListOptions{})
}