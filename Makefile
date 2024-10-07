DATE         ?= $(shell date +%FT%T%z)
REMOTE        = 127.0.0.1
TAG           = latest
RUNTIME       = docker

sources/operator/fetch-payload: ## Update tektoncd operator payloads and bundle manifests
	@./hack/operator-fetch-payload.sh
update/version/update-payload:
	@./hack/update-version.sh
	@./hack/operator-fetch-payload.sh

images: ${IMAGES} ## Build all images (except the bundle)
images/push: $(addsuffix /push,${IMAGES}) ## Build and push all images (except the bundle)
images/digest: $(addsuffix /digest,${IMAGES})

bundle: container/openshift-pipelines-operator-bundle/digest ## Build the operator bundle image
bundle/push: bundle container/openshift-pipelines-operator-bundle/push ## Build and push the bundle image
bundle/run: ## Run the pushed bundle (needs to run bundle/push target prior)
	operator-sdk run bundle -n openshift-operators ${REMOTE}/openshift-pipelines-operator-bundle:${TAG}
bundle/cleanup: ## Clean up a bundle installed in your cluster
	operator-sdk cleanup -n openshift-operators openshift-pipelines-operator-rh

index: bundle/push ## Create an index image for this bundle
	./hack/build-index.sh index/openshift-pipelines-operator-rh/package.yaml ${REMOTE} ${REMOTE}/openshift-pipelines-operator-bundle:${TAG} ${REMOTE}/openshift-pipelines-operator-index:${TAG} ${RUNTIME} ${DOCKER_BUILD_ARGS}
index/push: index ## Build and push the index image for this bundle
	${RUNTIME} ${DOCKER_BUILD_ARGS} push ${REMOTE}/openshift-pipelines-operator-index:${TAG}

FORCE:

.PHONY: help
help:
	@grep -hE '^[ a-zA-Z0-9/_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-32s\033[0m %s\n", $$1, $$2}'