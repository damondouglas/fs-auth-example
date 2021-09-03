help:		## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

web-server:	## Run on web-server device
	flutter run -d web-server --web-port=5000
