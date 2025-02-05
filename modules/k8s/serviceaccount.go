package k8s

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ServiceAccountHelper provides methods for managing Kubernetes ServiceAccounts
type ServiceAccountHelper interface {
	Get(ctx context.Context, namespace, name string) (*corev1.ServiceAccount, error)
	Create(ctx context.Context, namespace string, serviceAccount *corev1.ServiceAccount) (*corev1.ServiceAccount, error)
	Update(ctx context.Context, namespace string, serviceAccount *corev1.ServiceAccount) (*corev1.ServiceAccount, error)
	Delete(ctx context.Context, namespace, name string) error
	List(ctx context.Context, namespace string) (*corev1.ServiceAccountList, error)
}

type serviceAccountHelper struct {
	client *kubernetes.Clientset
}

// NewServiceAccountHelper creates a new instance of ServiceAccountHelper.
//
// Parameters:
//   - client: Kubernetes clientset for interacting with the cluster
//
// Returns:
//   - ServiceAccountHelper: Interface for managing Kubernetes ServiceAccounts
func NewServiceAccountHelper(client *kubernetes.Clientset) ServiceAccountHelper {
	return &serviceAccountHelper{
		client: client,
	}
}

// Get retrieves a service account by name from the specified namespace.
//
// Parameters:
//   - ctx: Context for the operation
//   - namespace: Kubernetes namespace
//   - name: Name of the service account
//
// Returns:
//   - *corev1.ServiceAccount: The retrieved service account
//   - error: nil if successful, error otherwise
func (sa *serviceAccountHelper) Get(ctx context.Context, namespace, name string) (*corev1.ServiceAccount, error) {
	return sa.client.CoreV1().ServiceAccounts(namespace).Get(ctx, name, metav1.GetOptions{})
}

// Create creates a new service account in the specified namespace.
//
// Parameters:
//   - ctx: Context for the operation
//   - namespace: Kubernetes namespace
//   - serviceAccount: ServiceAccount specification to create
//
// Returns:
//   - *corev1.ServiceAccount: The created service account
//   - error: nil if successful, error otherwise
func (sa *serviceAccountHelper) Create(ctx context.Context, namespace string, serviceAccount *corev1.ServiceAccount) (*corev1.ServiceAccount, error) {
	return sa.client.CoreV1().ServiceAccounts(namespace).Create(ctx, serviceAccount, metav1.CreateOptions{})
}

// Update updates an existing service account in the specified namespace.
//
// Parameters:
//   - ctx: Context for the operation
//   - namespace: Kubernetes namespace
//   - serviceAccount: Updated service account specification
//
// Returns:
//   - *corev1.ServiceAccount: The updated service account
//   - error: nil if successful, error otherwise
func (sa *serviceAccountHelper) Update(ctx context.Context, namespace string, serviceAccount *corev1.ServiceAccount) (*corev1.ServiceAccount, error) {
	return sa.client.CoreV1().ServiceAccounts(namespace).Update(ctx, serviceAccount, metav1.UpdateOptions{})
}

// Delete removes a service account from the specified namespace.
//
// Parameters:
//   - ctx: Context for the operation
//   - namespace: Kubernetes namespace
//   - name: Name of the service account to delete
//
// Returns:
//   - error: nil if successful, error otherwise
func (sa *serviceAccountHelper) Delete(ctx context.Context, namespace, name string) error {
	return sa.client.CoreV1().ServiceAccounts(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

// List retrieves all service accounts in the specified namespace.
//
// Parameters:
//   - ctx: Context for the operation
//   - namespace: Kubernetes namespace
//
// Returns:
//   - *corev1.ServiceAccountList: List of service accounts
//   - error: nil if successful, error otherwise
func (sa *serviceAccountHelper) List(ctx context.Context, namespace string) (*corev1.ServiceAccountList, error) {
	return sa.client.CoreV1().ServiceAccounts(namespace).List(ctx, metav1.ListOptions{})
}
