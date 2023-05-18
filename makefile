SHELL := /bin/bash

SERVICE_NAME 	:= k-island
# GO_FLAGS   		?= CGO_ENABLED=0 GOOS=linux GOARCH=amd64
OUTPUT_BIN 		?= bin/${SERVICE_NAME}
PACKAGE    		:= github.com/mboufous/$(SERVICE_NAME)
VERSION    		?= v1.0
IMAGE      		:= ${SERVICE_NAME}:${VERSION}

KIND_CLUSTER := dev-cluster

default: help

run:
	@go run main.go

test: # Run all tests
	go clean --testcache && go test ./... -count=1 --race

clean: ## Clean compiled service
	rm ${OUTPUT_BIN} || true

tidy: # tidy
	go mod tidy
	go mod vendor

.PHONY: build
build:  ## Builds the service
	@${GO_FLAGS} go build -o ${OUTPUT_BIN} \
	-ldflags "-w -s -X ${PACKAGE}/main.build=${VERSION}" \
	-a main.go

#--------------------------------
# Docker
img: clean build ## Build Docker image
	docker build \
		-f build/Dockerfile.dev \
		-t $(IMAGE) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.

#--------------------------------
# Kind

kind-up:  ## create a cluster and apply dev k8s configs
	kind create cluster --name $(KIND_CLUSTER)
	kubectl apply -f build/k8s/_test_cluster_setup
	
kind-dn: ## Delete cluster
	kind delete cluster --name $(KIND_CLUSTER)

kind-load: ## Load docker image into the cluster
	kind load docker-image $(SERVICE_NAME):$(VERSION) --name $(KIND_CLUSTER)

kind-apply: ## Apply k8s config files
	 kubectl apply -f build/k8s/kisland-pod
	 kubectl config set-context --current --namespace=dev-general

kind-logs: ## Show service logs from its pod
	kubectl logs -l app=kisland --all-containers=true -f --tail=100

kind-restart: ## Restart deployment
	kubectl rollout restart deployment kisland

kind-reload: all kind-load kind-restart ## Reload changes

kind-reload-apply: all kind-load kind-apply ## reapply k8s config changes

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":[^:]*?## "}; {printf "\033[38;5;69m%-30s\033[38;5;38m %s\033[0m\n", $$1, $$2}'
