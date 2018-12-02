.PHONY: build build-v1 server clean

YAML = docs/html/static/document.yaml

SRC = docs/source

build: build-v1

build-v1: .npminstall
	@echo Building dag openapi yaml

	@cat $(SRC)/introduction.yaml > $(YAML)
	@cat $(SRC)/definitions.yaml >> $(YAML)
	@node_modules/swagger-cli/bin/swagger-cli.js validate $(YAML)

	@echo Complete

.npminstall:
	@echo Getting dependencies using npm

	npm install swagger-cli@2.1.1
	touch $@

server:
	@echo the server is running at 8000

	make build
	cd docs/html && python2 -m SimpleHTTPServer

clean:
	@echo Cleaning

	rm -f .npminstall
	rm -rf node_modules
