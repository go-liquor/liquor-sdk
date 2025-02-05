# Kubernetes

## Contents
- [Enable](#enable)
- [Available Helpers](#available-helpers)
  - [Deployment](#deployment)
  - [Service](#service)
  - [Ingress](#ingress)
  - [ConfigMap](#configmap)
  - [Secret](#secret)
  - [ServiceAccount](#serviceaccount)
- [Usage Example](#usage-example)
- [Testing](#testing)

## Enable

```bash
liquor app enable k8s
# or
go get github.com/go-liquor/liquor-sdk/modules/k8s
```

In `cmd/app/main.go` add module:

```go
package main

import (
    "github.com/go-liquor/framework/internal/adapters/server/http"
    "github.com/go-liquor/framework/internal/app/services"
    "github.com/go-liquor/liquor-sdk/app"
    "github.com/go-liquor/liquor-sdk/modules/k8s" // this
)

func main() {
    app.NewApp(
        k8s.K8sModule, // this
    )
}
```

## Available Helpers

### Deployment

```go
type DeploymentHelper interface {
    Get(ctx context.Context, namespace, name string) (*appsv1.Deployment, error)
    Create(ctx context.Context, namespace string, deployment *appsv1.Deployment) (*appsv1.Deployment, error)
    Update(ctx context.Context, namespace string, deployment *appsv1.Deployment) (*appsv1.Deployment, error)
    Delete(ctx context.Context, namespace, name string) error
    List(ctx context.Context, namespace string) (*appsv1.DeploymentList, error)
    Scale(ctx context.Context, namespace, name string, replicas int32) error
}
```

### Service

```go
type ServiceHelper interface {
    Get(ctx context.Context, namespace, name string) (*corev1.Service, error)
    Create(ctx context.Context, namespace string, service *corev1.Service) (*corev1.Service, error)
    Update(ctx context.Context, namespace string, service *corev1.Service) (*corev1.Service, error)
    Delete(ctx context.Context, namespace, name string) error
    List(ctx context.Context, namespace string) (*corev1.ServiceList, error)
}
```


### Ingress

```go
type IngressHelper interface {
    Get(ctx context.Context, namespace, name string) (*networkingv1.Ingress, error)
    Create(ctx context.Context, namespace string, ingress *networkingv1.Ingress) (*networkingv1.Ingress, error)
    Update(ctx context.Context, namespace string, ingress *networkingv1.Ingress) (*networkingv1.Ingress, error)
    Delete(ctx context.Context, namespace, name string) error
    List(ctx context.Context, namespace string) (*networkingv1.IngressList, error)
}
```

### ConfigMap

```go
type ConfigMapHelper interface {
    Get(ctx context.Context, namespace, name string) (*corev1.ConfigMap, error)
    Create(ctx context.Context, namespace string, configMap *corev1.ConfigMap) (*corev1.ConfigMap, error)
    Update(ctx context.Context, namespace string, configMap *corev1.ConfigMap) (*corev1.ConfigMap, error)
    Delete(ctx context.Context, namespace, name string) error
    List(ctx context.Context, namespace string) (*corev1.ConfigMapList, error)
}
```

### Secret

```go
type SecretHelper interface {
    Get(ctx context.Context, namespace, name string) (*corev1.Secret, error)
    Create(ctx context.Context, namespace string, secret *corev1.Secret) (*corev1.Secret, error)
    Update(ctx context.Context, namespace string, secret *corev1.Secret) (*corev1.Secret, error)
    Delete(ctx context.Context, namespace, name string) error
    List(ctx context.Context, namespace string) (*corev1.SecretList, error)
}
```

### ServiceAccount

```go
type ServiceAccountHelper interface {
    Get(ctx context.Context, namespace, name string) (*corev1.ServiceAccount, error)
    Create(ctx context.Context, namespace string, serviceAccount *corev1.ServiceAccount) (*corev1.ServiceAccount, error)
    Update(ctx context.Context, namespace string, serviceAccount *corev1.ServiceAccount) (*corev1.ServiceAccount, error)
    Delete(ctx context.Context, namespace, name string) error
    List(ctx context.Context, namespace string) (*corev1.ServiceAccountList, error)
}
```

## Usage Example

In your service, you can inject and use any of these helpers:

```go
type Service struct {
    deployment     k8s.DeploymentHelper
    service        k8s.ServiceHelper
    ingress        k8s.IngressHelper
    configMap      k8s.ConfigMapHelper
    secret         k8s.SecretHelper
    serviceAccount k8s.ServiceAccountHelper
}

func NewService(
    deployment k8s.DeploymentHelper,
    service k8s.ServiceHelper,
    ingress k8s.IngressHelper,
    configMap k8s.ConfigMapHelper,
    secret k8s.SecretHelper,
    serviceAccount k8s.ServiceAccountHelper,
) *Service {
    return &Service{
        deployment:     deployment,
        service:        service,
        ingress:        ingress,
        configMap:      configMap,
        secret:        secret,
        serviceAccount: serviceAccount,
    }
}
```

## Testing

For testing purposes, you can use the generated mocks:

```go
func TestYourService(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockDeployment := k8s.NewMockDeploymentHelper(ctrl)
    mockService := k8s.NewMockServiceHelper(ctrl)
    mockIngress := k8s.NewMockIngressHelper(ctrl)
    mockConfigMap := k8s.NewMockConfigMapHelper(ctrl)
    mockSecret := k8s.NewMockSecretHelper(ctrl)
    mockServiceAccount := k8s.NewMockServiceAccountHelper(ctrl)

    // Configure expectations
    mockDeployment.EXPECT().
        Get(gomock.Any(), "default", "app").
        Return(&appsv1.Deployment{}, nil)

    // Initialize your service with mocks
    service := NewService(
        mockDeployment,
        mockService,
        mockIngress,
        mockConfigMap,
        mockSecret,
        mockServiceAccount,
    )

    // Run your tests
    result, err := service.YourMethod()
    require.NoError(t, err)
}
```