# debug
run:
	go run api/services/sales/main.go | go run api/tooling/logfmt/main.go

# Define dependencies

GOLANG          := golang:1.25.1
ALPINE          := alpine:3.22
KIND            := kindest/node:v1.34.0
POSTGRES        := postgres:18.0
GRAFANA         := grafana/grafana:12.2.0
PROMETHEUS      := prom/prometheus:v3.6.0
TEMPO           := grafana/tempo:2.8.1
LOKI            := grafana/loki:3.5.0
PROMTAIL        := grafana/promtail:3.5.0

KIND_CLUSTER    := atlas-starter-cluster
NAMESPACE       := sales-system
SALES_APP       := sales
AUTH_APP        := auth
BASE_IMAGE_NAME := localhost/atlas
VERSION         := 0.0.1
SALES_IMAGE     := $(BASE_IMAGE_NAME)/$(SALES_APP):$(VERSION)
METRICS_IMAGE   := $(BASE_IMAGE_NAME)/metrics:$(VERSION)
AUTH_IMAGE      := $(BASE_IMAGE_NAME)/$(AUTH_APP):$(VERSION)

# VERSION       := "0.0.1-$(shell git rev-parse --short HEAD)"

# ==============================================================================
# building containers

build: sales

sales:
	docker build \
		-f zarf/docker/dockerfile.sales \
		-t $(SALES_IMAGE) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
		.
# ==============================================================================
# module support
tidy:
	go mod tidy
	go mod vendor

# ==============================================================================
# running from within k8s/kind
dev-up:
	kind create cluster \
		--image $(KIND) \
		--name $(KIND_CLUSTER) \
		--config zarf/k8s/dev/kind-config.yaml

	kubectl wait --timeout=120s --namespace=local-path-storage --for=condition=Available deployment/local-path-provisioner

dev-down:
	kind delete cluster --name $(KIND_CLUSTER)

dev-status-all:
	kubectl get nodes -o wide
	kubectl get svc -o wide
	kubectl get pods -o wide --watch --all-namespaces

dev-status:
	watch -n 2 kubectl get pods -o wide --all-namespaces
