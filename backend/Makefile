PROJECT=$$(gcloud config get-value project)

help:	## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

build:	## Build container image
	export KO_DOCKER_REPO=gcr.io/$(PROJECT); ko publish ./cmd/service
