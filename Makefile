.PHONY: all

DOCKER_URL := hibri/sample-k8s-controller
TAG := latest

print-%: ; @echo $*=$($*)

default: build

codegen:
	hack/update-codegen.sh
	dep ensure

test:
	hack/test.sh

build: test
	hack/build.sh

build-for-docker: test
	hack/build-for-docker.sh

package: build-for-docker
	docker build -t "$(DOCKER_URL):$(TAG)" -f Dockerfile .

publish: package
	docker push "$(DOCKER_URL):$(TAG)"

deploy-crd:
	kubectl apply -f artifacts/crd.yaml
	kubectl apply -f artifacts/crd-validation.yaml

undeploy-crd:
	kubectl delete -f artifacts/crd-validation.yaml --ignore-not-found=true
	kubectl delete -f artifacts/crd.yaml --ignore-not-found=true
	
deploy-cr:
	kubectl apply -f artifacts/cr.yaml

undeploy-cr:
	kubectl delete -f artifacts/cr.yaml --ignore-not-found=true

deploy-controller:
	kubectl apply -f artifacts/controller.yaml

undeploy-controller:
	kubectl delete -f artifacts/controller.yaml --ignore-not-found=true

run-inside-cluster: deploy-crd deploy-cr deploy-controller
	
run-outside-cluster: deploy-crd deploy-cr
	sample-k8s-controller -logtostderr=true -v=2 -stderrthreshold=INFO

cleanup: undeploy-controller undeploy-cr undeploy-crd 
	kubectl delete deployment example-foo --ignore-not-found=true

#call "make cleanup" before running demo 
demo-1: 
	clear
	kubectl get pods
	kubectl apply -f artifacts/crd.yaml
	kubectl apply -f artifacts/crd-validation.yaml
	kubectl apply -f artifacts/controller.yaml
	sleep 5
	kubectl get pods
	sleep 5 # so what happened
	
	clear
	kubectl get pods
	kubectl get deployments
	kubectl apply -f artifacts/cr.yaml
	sleep 5
	kubectl get pods
	kubectl get deployments
	sleep 5 #so what happened

	clear
	kubectl get pods
	kubectl get deployments
	hack/delete-sample-foo-pods.sh
	sleep 1
	kubectl get pods
	kubectl get deployments
	sleep 5 #so what happened

help: ## this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

