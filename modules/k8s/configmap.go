package k8s

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ConfigMapHelper interface {
	Get(ctx context.Context, namespace, name string) (*corev1.ConfigMap, error)
	Create(ctx context.Context, namespace string, configMap *corev1.ConfigMap) (*corev1.ConfigMap, error)
	Update(ctx context.Context, namespace string, configMap *corev1.ConfigMap) (*corev1.ConfigMap, error)
	Delete(ctx context.Context, namespace, name string) error
	List(ctx context.Context, namespace string) (*corev1.ConfigMapList, error)
}

type configMapHelper struct {
	client *kubernetes.Clientset
}

func NewConfigMapHelper(client *kubernetes.Clientset) ConfigMapHelper {
	return &configMapHelper{
		client: client,
	}
}

func (c *configMapHelper) Get(ctx context.Context, namespace, name string) (*corev1.ConfigMap, error) {
	return c.client.CoreV1().ConfigMaps(namespace).Get(ctx, name, metav1.GetOptions{})
}

func (c *configMapHelper) Create(ctx context.Context, namespace string, configMap *corev1.ConfigMap) (*corev1.ConfigMap, error) {
	return c.client.CoreV1().ConfigMaps(namespace).Create(ctx, configMap, metav1.CreateOptions{})
}

func (c *configMapHelper) Update(ctx context.Context, namespace string, configMap *corev1.ConfigMap) (*corev1.ConfigMap, error) {
	return c.client.CoreV1().ConfigMaps(namespace).Update(ctx, configMap, metav1.UpdateOptions{})
}

func (c *configMapHelper) Delete(ctx context.Context, namespace, name string) error {
	return c.client.CoreV1().ConfigMaps(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

func (c *configMapHelper) List(ctx context.Context, namespace string) (*corev1.ConfigMapList, error) {
	return c.client.CoreV1().ConfigMaps(namespace).List(ctx, metav1.ListOptions{})
}
