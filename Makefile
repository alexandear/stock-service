MAKEFILE_PATH := $(abspath $(dir $(abspath $(lastword $(MAKEFILE_LIST)))))
PATH := $(MAKEFILE_PATH):$(PATH)

export GOBIN := $(MAKEFILE_PATH)/bin

PATH := $(GOBIN):$(PATH)

.PHONY: all
all: clean format build lint test docker

.PHONY: clean
clean:
	@echo clean
	@go clean

.PHONY: build
build:
	@echo build
	@go build -o $(GOBIN)/stock-service

.PHONY: test
test:
	@echo test
	@go test -count=1 -v ./...

.PHONY: lint
lint:
	@echo lint
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint
	@$(GOBIN)/golangci-lint run

.PHONY: format
format:
	@echo format
	@go fmt $(PKGS)

.PHONY: generate
generate: mock grpc

.PHONY: mock
mock:
	@echo mock
	@go install github.com/golang/mock/mockgen
	@go generate ./...

.PHONY: grpc
grpc:
	@echo grpc
	@go install github.com/golang/protobuf/protoc-gen-go

	@-rm -rf $(MAKEFILE_PATH)/internal/grpcapi
	@mkdir -p $(MAKEFILE_PATH)/internal/grpcapi

	@protoc -I $(MAKEFILE_PATH)/api $(MAKEFILE_PATH)/api/stock.proto \
		--go_out=plugins=grpc:$(MAKEFILE_PATH)/internal/grpcapi

IMAGE = stock-service

.PHONY: docker
docker:
	@echo docker
	@docker build -t $(IMAGE) -f Dockerfile .

.PHONY: docker-run
docker-run:
	@echo docker-run
	@docker run --rm -p 9090:9090 $(IMAGE)

IMAGE_DEV = stock-service-dev

.PHONY: docker-dev
docker-dev:
	@echo docker-dev
	@docker build -t $(IMAGE_DEV) -f Dockerfile.build .
