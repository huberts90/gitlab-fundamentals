DOCKER_IMAGE := gitlab-fundamentals

ifdef VERSION
	DOCKER_VERSION = $(VERSION)
else
	DOCKER_VERSION := 1.0.0
endif

.PHONY: test
test: ## test using your local toolchain
	go test ./internal/...

.PHONY: docker-build
docker-build: ## build using the official toolchain
	docker build \
		-f Dockerfile \
		-t $(DOCKER_IMAGE):$(DOCKER_VERSION)-build \
		--target=build \
		.

.PHONY: docker-test
docker-test: ## build using the official toolchain
	docker build \
		-f Dockerfile \
		-t $(DOCKER_IMAGE):$(DOCKER_VERSION)-test \
		--target=test \
		.

.PHONY: docker-release
docker-release: docker-build ## copy release from docker

	##  `docker create` command creates a writeable container layer but never started
	docker create --name $(DOCKER_IMAGE)_$(DOCKER_VERSION)-build $(DOCKER_IMAGE):$(DOCKER_VERSION)-build
	docker cp $(DOCKER_IMAGE)_$(DOCKER_VERSION)-build:/usr/src/app/release .
	docker rm -f $(DOCKER_IMAGE)_$(DOCKER_VERSION)-build