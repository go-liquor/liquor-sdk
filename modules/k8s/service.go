package k8s

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ServiceHelper provides methods for managing Kubernetes Services
type ServiceHelper interface {
	Get(ctx context.Context, namespace, name string) (*corev1.Service, error)
	Create(ctx context.Context, namespace string, service *corev1.Service) (*corev1.Service, error)
	Update(ctx context.Context, namespace string, service *corev1.Service) (*corev1.Service, error)
	Delete(ctx context.Context, namespace, name string) error
	List(ctx context.Context, namespace string) (*corev1.ServiceList, error)
}

type serviceHelper struct {
	client *kubernetes.Clientset
}

// NewServiceHelper creates a new instance of ServiceHelper.
//
// Parameters:
//   - client: Kubernetes clientset for interacting with the cluster
//
// Returns:
//   - ServiceHelper: Interface for managing Kubernetes Services
func NewServiceHelper(client *kubernetes.Clientset) ServiceHelper {
	return &serviceHelper{
		client: client,
	}
}

// Get retrieves a service by name from the specified namespace.
//
// Parameters:
//   - ctx: Context for the operation
//   - namespace: Kubernetes namespace
//   - name: Name of the service
//
// Returns:
//   - *corev1.Service: The retrieved service
//   - error: nil if successful, error otherwise
func (s *serviceHelper) Get(ctx context.Context, namespace, name string) (*corev1.Service, error) {
	return s.client.CoreV1().Services(namespace).Get(ctx, name, metav1.GetOptions{})
}

// Create creates a new service in the specified namespace.
//
// Parameters:
//   - ctx: Context for the operation
//   - namespace: Kubernetes namespace
//   - service: Service specification to create
//
// Returns:
//   - *corev1.Service: The created service
//   - error: nil if successful, error otherwise
func (s *serviceHelper) Create(ctx context.Context, namespace string, service *corev1.Service) (*corev1.Service, error) {
	return s.client.CoreV1().Services(namespace).Create(ctx, service, metav1.CreateOptions{})
}

// Update updates an existing service in the specified namespace.
//
// Parameters:
//   - ctx: Context for the operation
//   - namespace: Kubernetes namespace
//   - service: Updated service specification
//
// Returns:
//   - *corev1.Service: The updated service
//   - error: nil if successful, error otherwise
func (s *serviceHelper) Update(ctx context.Context, namespace string, service *corev1.Service) (*corev1.Service, error) {
	return s.client.CoreV1().Services(namespace).Update(ctx, service, metav1.UpdateOptions{})
}

// Delete removes a service from the specified namespace.
//
// Parameters:
//   - ctx: Context for the operation
//   - namespace: Kubernetes namespace
//   - name: Name of the service to delete
//
// Returns:
//   - error: nil if successful, error otherwise
func (s *serviceHelper) Delete(ctx context.Context, namespace, name string) error {
	return s.client.CoreV1().Services(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

// List retrieves all services in the specified namespace.
//
// Parameters:
//   - ctx: Context for the operation
//   - namespace: Kubernetes namespace
//
// Returns:
//   - *corev1.ServiceList: List of services
//   - error: nil if successful, error otherwise
func (s *serviceHelper) List(ctx context.Context, namespace string) (*corev1.ServiceList, error) {
	return s.client.CoreV1().Services(namespace).List(ctx, metav1.ListOptions{})
}
