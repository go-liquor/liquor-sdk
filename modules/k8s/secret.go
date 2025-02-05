package k8s

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// SecretHelper provides methods for managing Kubernetes Secrets
type SecretHelper interface {
	Get(ctx context.Context, namespace, name string) (*corev1.Secret, error)
	Create(ctx context.Context, namespace string, secret *corev1.Secret) (*corev1.Secret, error)
	Update(ctx context.Context, namespace string, secret *corev1.Secret) (*corev1.Secret, error)
	Delete(ctx context.Context, namespace, name string) error
	List(ctx context.Context, namespace string) (*corev1.SecretList, error)
}

type secretHelper struct {
	client *kubernetes.Clientset
}

// NewSecretHelper creates a new instance of SecretHelper.
//
// Parameters:
//   - client: Kubernetes clientset for interacting with the cluster
//
// Returns:
//   - SecretHelper: Interface for managing Kubernetes Secrets
func NewSecretHelper(client *kubernetes.Clientset) SecretHelper {
	return &secretHelper{
		client: client,
	}
}

// Get retrieves a secret by name from the specified namespace.
//
// Parameters:
//   - ctx: Context for the operation
//   - namespace: Kubernetes namespace
//   - name: Name of the secret
//
// Returns:
//   - *corev1.Secret: The retrieved secret
//   - error: nil if successful, error otherwise
func (s *secretHelper) Get(ctx context.Context, namespace, name string) (*corev1.Secret, error) {
	return s.client.CoreV1().Secrets(namespace).Get(ctx, name, metav1.GetOptions{})
}

// Create creates a new secret in the specified namespace.
//
// Parameters:
//   - ctx: Context for the operation
//   - namespace: Kubernetes namespace
//   - secret: Secret specification to create
//
// Returns:
//   - *corev1.Secret: The created secret
//   - error: nil if successful, error otherwise
func (s *secretHelper) Create(ctx context.Context, namespace string, secret *corev1.Secret) (*corev1.Secret, error) {
	return s.client.CoreV1().Secrets(namespace).Create(ctx, secret, metav1.CreateOptions{})
}

// Update updates an existing secret in the specified namespace.
//
// Parameters:
//   - ctx: Context for the operation
//   - namespace: Kubernetes namespace
//   - secret: Updated secret specification
//
// Returns:
//   - *corev1.Secret: The updated secret
//   - error: nil if successful, error otherwise
func (s *secretHelper) Update(ctx context.Context, namespace string, secret *corev1.Secret) (*corev1.Secret, error) {
	return s.client.CoreV1().Secrets(namespace).Update(ctx, secret, metav1.UpdateOptions{})
}

// Delete removes a secret from the specified namespace.
//
// Parameters:
//   - ctx: Context for the operation
//   - namespace: Kubernetes namespace
//   - name: Name of the secret to delete
//
// Returns:
//   - error: nil if successful, error otherwise
func (s *secretHelper) Delete(ctx context.Context, namespace, name string) error {
	return s.client.CoreV1().Secrets(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

// List retrieves all secrets in the specified namespace.
//
// Parameters:
//   - ctx: Context for the operation
//   - namespace: Kubernetes namespace
//
// Returns:
//   - *corev1.SecretList: List of secrets
//   - error: nil if successful, error otherwise
func (s *secretHelper) List(ctx context.Context, namespace string) (*corev1.SecretList, error) {
	return s.client.CoreV1().Secrets(namespace).List(ctx, metav1.ListOptions{})
}