package k8s

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// DeploymentHelper provides methods for managing Kubernetes Deployments
type DeploymentHelper interface {
	Get(ctx context.Context, namespace, name string) (*appsv1.Deployment, error)
	Create(ctx context.Context, namespace string, deployment *appsv1.Deployment) (*appsv1.Deployment, error)
	Update(ctx context.Context, namespace string, deployment *appsv1.Deployment) (*appsv1.Deployment, error)
	Delete(ctx context.Context, namespace, name string) error
	List(ctx context.Context, namespace string) (*appsv1.DeploymentList, error)
	Scale(ctx context.Context, namespace, name string, replicas int32) error
}

type deploymentHelper struct {
	client *kubernetes.Clientset
}

// NewDeploymentHelper creates a new instance of DeploymentHelper.
//
// Parameters:
//   - client: Kubernetes clientset for interacting with the cluster
//
// Returns:
//   - DeploymentHelper: Interface for managing Kubernetes Deployments
func NewDeploymentHelper(client *kubernetes.Clientset) DeploymentHelper {
	return &deploymentHelper{
		client: client,
	}
}

// Get retrieves a deployment by name from the specified namespace.
//
// Parameters:
//   - ctx: Context for the operation
//   - namespace: Kubernetes namespace
//   - name: Name of the deployment
//
// Returns:
//   - *appsv1.Deployment: The retrieved deployment
//   - error: nil if successful, error otherwise
func (d *deploymentHelper) Get(ctx context.Context, namespace, name string) (*appsv1.Deployment, error) {
	return d.client.AppsV1().Deployments(namespace).Get(ctx, name, metav1.GetOptions{})
}

// Create creates a new deployment in the specified namespace.
//
// Parameters:
//   - ctx: Context for the operation
//   - namespace: Kubernetes namespace
//   - deployment: Deployment specification to create
//
// Returns:
//   - *appsv1.Deployment: The created deployment
//   - error: nil if successful, error otherwise
func (d *deploymentHelper) Create(ctx context.Context, namespace string, deployment *appsv1.Deployment) (*appsv1.Deployment, error) {
	return d.client.AppsV1().Deployments(namespace).Create(ctx, deployment, metav1.CreateOptions{})
}

// Update updates an existing deployment in the specified namespace.
//
// Parameters:
//   - ctx: Context for the operation
//   - namespace: Kubernetes namespace
//   - deployment: Updated deployment specification
//
// Returns:
//   - *appsv1.Deployment: The updated deployment
//   - error: nil if successful, error otherwise
func (d *deploymentHelper) Update(ctx context.Context, namespace string, deployment *appsv1.Deployment) (*appsv1.Deployment, error) {
	return d.client.AppsV1().Deployments(namespace).Update(ctx, deployment, metav1.UpdateOptions{})
}

// Delete removes a deployment from the specified namespace.
//
// Parameters:
//   - ctx: Context for the operation
//   - namespace: Kubernetes namespace
//   - name: Name of the deployment to delete
//
// Returns:
//   - error: nil if successful, error otherwise
func (d *deploymentHelper) Delete(ctx context.Context, namespace, name string) error {
	return d.client.AppsV1().Deployments(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

// List retrieves all deployments in the specified namespace.
//
// Parameters:
//   - ctx: Context for the operation
//   - namespace: Kubernetes namespace
//
// Returns:
//   - *appsv1.DeploymentList: List of deployments
//   - error: nil if successful, error otherwise
func (d *deploymentHelper) List(ctx context.Context, namespace string) (*appsv1.DeploymentList, error) {
	return d.client.AppsV1().Deployments(namespace).List(ctx, metav1.ListOptions{})
}

// Scale adjusts the number of replicas for a deployment.
//
// Parameters:
//   - ctx: Context for the operation
//   - namespace: Kubernetes namespace
//   - name: Name of the deployment
//   - replicas: Desired number of replicas
//
// Returns:
//   - error: nil if successful, error otherwise
func (d *deploymentHelper) Scale(ctx context.Context, namespace, name string, replicas int32) error {
	deployment, err := d.Get(ctx, namespace, name)
	if err != nil {
		return err
	}

	deployment.Spec.Replicas = &replicas
	_, err = d.Update(ctx, namespace, deployment)
	return err
}
