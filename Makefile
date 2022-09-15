.PHONY: generate

DOCKER_IMAGE_NAME=api-builder
ifeq ($(OS),Windows_NT)
PWD=${CURDIR}
endif

fast-generate:
	@buf generate --path "api/proto"

docker-build:
	@echo "Building docker image '${DOCKER_IMAGE_NAME}'"
	@docker build -t ${DOCKER_IMAGE_NAME} -f ./tools/docker/Dockerfile .
	@echo "The docker image building finished successfully"

docker-generate:
	@echo "Generating using docker image '${DOCKER_IMAGE_NAME}'"
	@docker run -it -v "${PWD}":/src "${DOCKER_IMAGE_NAME}"

install-buf:
	@go install github.com/bufbuild/buf/cmd/buf@v1.4.0

install-go-deps:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.0
	@go install github.com/twitchtv/twirp/protoc-gen-twirp@v8.1.2
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.10.3
	@go install github.com/envoyproxy/protoc-gen-validate@v0.6.7
