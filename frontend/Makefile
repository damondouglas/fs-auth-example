help:		## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

web-server:	## Run on web-server device
	flutter run -d web-server --web-port=5000

build-web:	## Build the web application
	flutter build web

deploy-web:	## Deploy the web application
	firebase deploy
