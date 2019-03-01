version = 1.0.6
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
.PHONY: run
run: build
	@docker run -it -p 8080:8080 $(registry)/$(build)
