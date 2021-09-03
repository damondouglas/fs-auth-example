PROJECT=$$(gcloud config get-value project)

help:	## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

run: 	## Run the service
	go run ./cmd/service --local --verbose

build-deploy-tcp: ## Deploy the TCP service
	gcloud run deploy counter-tcp \
		--image=$$(export KO_DOCKER_REPO=gcr.io/$(PROJECT); ko publish ./cmd/service) \
		--port=50051 \
		--set-env-vars=TCP_PORT=50051,HTTP_PORT=8080,PROJECT=$(PROJECT) \
		--allow-unauthenticated \
		--region=us-central1

build-deploy-http: ## Deploy the HTTP service
	gcloud run deploy counter-http \
		--image=$$(export KO_DOCKER_REPO=gcr.io/$(PROJECT); ko publish ./cmd/service) \
		--port=8080 \
		--set-env-vars=TCP_PORT=50051,HTTP_PORT=8080,PROJECT=$(PROJECT) \
		--allow-unauthenticated \
		--region=us-central1
