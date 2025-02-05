package k8s

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// NewClient creates a new Kubernetes clientset using the provided configuration.
// It will log a fatal error if the client creation fails.
//
// Parameters:
//   - logger: A zap logger instance for error logging
//   - config: The Kubernetes REST configuration
//
// Returns:
//   - *kubernetes.Clientset: A Kubernetes client for interacting with the cluster
func NewClient(logger *zap.Logger, config *rest.Config) *kubernetes.Clientset {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Fatal("error to create client For Config", zap.Error(err))
	}
	return clientset
}

// NewRestConfig creates a new Kubernetes REST configuration.
// It first attempts to use in-cluster configuration, and if that fails,
// falls back to local configuration from ~/.kube/config.
//
// Parameters:
//   - logger: A zap logger instance for error logging
//
// Returns:
//   - *rest.Config: The Kubernetes REST configuration
func NewRestConfig(logger *zap.Logger) *rest.Config {
	config, err := rest.InClusterConfig()
	if err != nil {
		logger.Warn("Unable to use kubernetes config through cluster, trying local configuration", zap.Error(err))
		home, _ := os.UserHomeDir()
		config, err = clientcmd.BuildConfigFromFlags("", filepath.Join(home, ".kube", "config"))
		if err != nil {
			logger.Fatal("error to create config", zap.Error(err))
		}
	}
	return config
}

// NewDynamicClient creates a new Kubernetes dynamic client using the provided configuration.
// Dynamic client allows interaction with CRDs and other dynamic resources.
// It will log a fatal error if the client creation fails.
//
// Parameters:
//   - logger: A zap logger instance for error logging
//   - config: The Kubernetes REST configuration
//
// Returns:
//   - *dynamic.DynamicClient: A dynamic client for interacting with the cluster
func NewDynamicClient(logger *zap.Logger, config *rest.Config) *dynamic.DynamicClient {
	clientset, err := dynamic.NewForConfig(config)
	if err != nil {
		logger.Fatal("error to create DynamicClient For Client", zap.Error(err))
	}
	return clientset
}
