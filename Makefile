version = 1.0.1
image = cowsay
registry = guygrigsby
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
