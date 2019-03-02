version = 2.0.12
image = cowsay
registry = docker.io/guygrigsby
#registry = 819820547151.dkr.ecr.us-west-2.amazonaws.com
build = $(image):$(version)

.PHONY: build
build:
	@echo "Building $(build)..."
	@docker build --rm=true --no-cache=true --pull=true -t $(build) .
	@docker tag $(build) $(registry)/$(build)

.PHONY: release
release: build
	@echo "Releasing $(build)..."
	@docker push $(registry)/$(build)
.PHONY: run
run: build
	@docker run -it -p 8080:80 $(registry)/$(build)
.PHONY: test
test:
	go test ./...
