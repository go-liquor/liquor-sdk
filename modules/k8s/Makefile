.PHONY: genmocks
genmocks:
	mockgen -source=configmap.go -package=k8s -destination=configmap_mock.go
	mockgen -source=deployment.go -package=k8s -destination=deployment_mock.go
	mockgen -source=ingress.go -package=k8s -destination=ingress_mock.go
	mockgen -source=secret.go -package=k8s -destination=secret_mock.go
	mockgen -source=service.go -package=k8s -destination=service_mock.go
	mockgen -source=serviceaccount.go -package=k8s -destination=serviceaccount_mock.go