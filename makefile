# debug
run:
	go run api/services/sales/main.go | go run api/tooling/logfmt/main.go

help:
	go run api/services/sales/main.go --help

version:
	go run api/services/sales/main.go --version

# temp
curl-test-api:
	curl -il -H GET http://localhost:3000/test

curl-liveness:
	curl -il -H GET http://localhost:3000/liveness

curl-readiness:
	curl -il -H GET http://localhost:3000/readiness

curl-test-error:
	curl -il -H GET http://localhost:3000/testerror

curl-test-panic:
	curl -il -H GET http://localhost:3000/testpanic

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

# ENV (IMPORTANT: Should be secret in real projects!)
DB_USER         := postgres
DB_PASS         := postgres

# VERSION       := "0.0.1-$(shell git rev-parse --short HEAD)"

# ==============================================================================
# Install dependencies

dev-gotooling:
	go install github.com/divan/expvarmon@latest
	go install github.com/rakyll/hey@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install golang.org/x/tools/cmd/goimports@latest

dev-docker:
	docker pull $(GOLANG) & \
	docker pull $(ALPINE) & \
	docker pull $(KIND) & \
	docker pull $(POSTGRES) & \
	docker pull $(GRAFANA) & \
	docker pull $(PROMETHEUS) & \
	docker pull $(TEMPO) & \
	docker pull $(LOKI) & \
	docker pull $(PROMTAIL) & \
	wait;


# ==============================================================================
# building containers

build: sales auth

sales:
	docker build \
		-f zarf/docker/dockerfile.sales \
		-t $(SALES_IMAGE) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
		.

auth:
	docker build \
		-f zarf/docker/dockerfile.auth \
		-t $(AUTH_IMAGE) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
		.

# ==============================================================================
# Administration

dev-db:
	kubectl port-forward svc/database-service -n sales-system 5433:5432 & sleep 1 && pgcli postgresql://$(DB_USER):$(DB_PASS)@localhost:5433/postgres

dev-db-stop:
	-pkill -f "port-forward svc/database-service"
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

dev-recreate:
	kind delete cluster --name $(KIND_CLUSTER) || true
	make dev-up dev-load dev-apply

dev-status-all:
	kubectl get nodes -o wide
	kubectl get svc -o wide
	kubectl get pods -o wide --watch --all-namespaces

dev-status:
	watch -n 2 kubectl get pods -o wide --all-namespaces

# ==============================================================================

dev-load-db:
	kind load docker-image $(POSTGRES) --name $(KIND_CLUSTER)

dev-load:
	kind load docker-image $(SALES_IMAGE) --name $(KIND_CLUSTER)
	kind load docker-image $(AUTH_IMAGE) --name $(KIND_CLUSTER) 

dev-apply:

	kustomize build zarf/k8s/dev/database | kubectl apply -f -
	kubectl rollout status --namespace=$(NAMESPACE) --watch --timeout=120s sts/database
	
	kustomize build zarf/k8s/dev/auth | kubectl apply -f -
	kubectl wait pods --namespace=$(NAMESPACE) --selector app=$(AUTH_APP) --timeout=120s --for=condition=Ready

	kustomize build zarf/k8s/dev/sales | kubectl apply -f -
	kubectl wait pods --namespace=$(NAMESPACE) --selector app=$(SALES_APP) --timeout=120s --for=condition=Ready
	
dev-restart:
	kubectl rollout restart deployment $(AUTH_APP) --namespace=$(NAMESPACE)
	kubectl rollout restart deployment $(SALES_APP) --namespace=$(NAMESPACE)

dev-run: build dev-up dev-load-db dev-load dev-apply

dev-update: build dev-load dev-restart

dev-update-apply: build dev-load dev-apply

dev-logs-sales:
	kubectl logs --namespace=$(NAMESPACE) -l app=$(SALES_APP) --all-containers=true -f --tail=100 --max-log-requests=6 | go run api/tooling/logfmt/main.go -service=$(SALES_APP)

dev-logs-sales-verbose:
	kubectl logs --namespace=$(NAMESPACE) -l app=$(SALES_APP) --all-containers=true -f --tail=100 --max-log-requests=6

dev-logs-auth:
	kubectl logs --namespace=$(NAMESPACE) -l app=$(AUTH_APP) --all-containers=true -f --tail=100 --max-log-requests=6 | go run api/tooling/logfmt/main.go -service=$(AUTH_APP)

dev-logs-auth-verbose:
	kubectl logs --namespace=$(NAMESPACE) -l app=$(AUTH_APP) --all-containers=true -f --tail=100 --max-log-requests=6

dev-logs-init:
	kubectl logs --namespace=$(NAMESPACE) -l app=$(SALES_APP) -f --tail=100 -c init-migrate-seed

dev-get-svc:
	kubectl get svc -n $(NAMESPACE) -o wide

dev-forward-port:
	kubectl port-forward -n $(NAMESPACE) svc/$(NAMESPACE) 3000:3000

dev-postgres-forward-port:
	kubectl port-forward svc/database-service -n sales-system 5433:5432

# ==============================================================================

dev-describe-deployment:
	kubectl describe deployment --namespace=$(NAMESPACE) $(SALES_APP)

dev-describe-sales:
	kubectl describe pod --namespace=$(NAMESPACE) -l app=$(SALES_APP)

dev-describe-auth:
	kubectl describe pod --namespace=$(NAMESPACE) -l app=$(AUTH_APP)

# ==============================================================================
# Metrics and Tracing

metrics:
	expvarmon -ports="localhost:3010" -vars="build,requests,goroutines,errors,panics,mem:memstats.HeapAlloc,mem:memstats.HeapSys,mem:memstats.Sys"

statsviz:
	open http://localhost:3010/debug/statsviz

# ==============================================================================
# Running tests within the local computer

test-r:
	CGO_ENABLED=1 go test -race -count=1 ./...

test-only:
	CGO_ENABLED=0 go test -count=1 ./...

lint:
	CGO_ENABLED=0 go vet ./...
	staticcheck -checks=all ./...

vuln-check:
	govulncheck ./...

test: test-only lint vuln-check

test-race: test-r lint vuln-check


# ==============================================================================
# Misc

# RSA Keys
# 	To generate a private/public key PEM file.
# 	$ openssl genpkey -algorithm RSA -out private.pem -pkeyopt rsa_keygen_bits:2048
# 	$ openssl rsa -pubout -in private.pem -out public.pem
# 	$ ./admin genkey

admin-pvt-key:
	go run api/tooling/admin/main.go

# admin token:
# export TOKEN=eyJhbGciOiJSUzI1NiIsImtpZCI6IjdmN2I4OTkzLWYyZDYtNDJmOS1hMGU4LTFjMmVmMTZjNTRjZCIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzZXJ2aWNlIHByb2plY3QiLCJzdWIiOiIzMTUyNDIwZC1jMDVlLTQxOWUtYTQyYS03ZTVlYjExNTFjYWMiLCJleHAiOjE3OTQ0NjI0ODAsImlhdCI6MTc2MjkyNjQ4MCwiUm9sZXMiOlsiQURNSU4iXX0.pgp8PqGyhOXISXC_Qv4ydS7WbLYhkN6wr05Hqc6GNCsADchF4gZn1NcLkMJ_Z9oHJBYYNq3Pgp9ng5FE7jboFqKfi-Dnm137raxYo3vLl4A9mqm8zaFIPEh9IlgnzvVU1usCjbOzqrX1c5w8C4IP45XkbWYPK2KrjkTYgcsGGeUyi0lvBra3o-LfbvmcIZ3hN7dK13J-gAgEp2LMGvbF0J65HBwU9ZOz5lsDeQHKDCbBWH1PRMAP5uyNBYEruLPgH0x_cfMP3cZj7dHJd91QLo8Co92mvWW6u7P35N1BgEFdR1fPCwSBj_b6ePvZgQeAls-kUDOpFHWfPQyPAjlEtg
# user token:
# export TOKEN=eyJhbGciOiJSUzI1NiIsImtpZCI6IjdmN2I4OTkzLWYyZDYtNDJmOS1hMGU4LTFjMmVmMTZjNTRjZCIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzZXJ2aWNlIHByb2plY3QiLCJzdWIiOiIzMTUyNDIwZC1jMDVlLTQxOWUtYTQyYS03ZTVlYjExNTFjYWMiLCJleHAiOjE3OTQ0NjI2NTUsImlhdCI6MTc2MjkyNjY1NSwiUm9sZXMiOlsiVVNFUiJdfQ.rsn-FrVhwkA4FnYf4cD4HlqA_JP62XWKTFeMhElIDdRvmxgN-NObgbJjk_yWhxmrg8ID9Q_OkEI0D2gRM60IuOi_k6JWwxVpgjwlT_AI_iNpVrMA42NGa5MxYwfltTuVanbtZvd0D_pi-UfksnjdHMQqVPkDzZK6qWtyD_FvsLYFY7OUnydDwG1ay3hi5FU1CO9c63KF1T_X2l7kOSIvpDyi72sxqYBrK-wCwGyK_4H2hMH6gcdKtb6nG3-qtNKGpJWflFszGPOTtDMSqbiBGrbIkXC0NKB1KvEHsitjmVCV3lfk3DdhRuxKtgCFqxZOD59k6w9I258nxybvyXOAgw

users:
	curl -il \
	-H "Authorization: Bearer ${TOKEN}" "http://localhost:3000/users?page=1&rows=2"

curl-auth:
	curl -il \
	-H "Authorization: Bearer ${TOKEN}" "http://localhost:3000/testauth"

token:
	curl -il \
	--user "admin@example.com:gophers" http://localhost:6000/auth/token/7f7b8993-f2d6-42f9-a0e8-1c2ef16c54cd

curl-auth2:
	curl -il \
	-H "Authorization: Bearer ${TOKEN}" "http://localhost:6000/auth/authenticate"