# Name of this Docker-Container
NAME = padel-field

# Relative path to bin (scripts) directory (you may need to add/remove ".." paths)
DIR_API = /api
DIR_UI = /ui

.DEFAULT_GOAL := help
help: ## Print the list of makefile targets
	@grep -hE '^[%a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "%-18s %s\n", $$1, $$2}'

#
#  BUILD AND RUN
#

build: ## Build the docker image
	docker build . -t $(NAME) -f 'Dockerfile'

run: # start the engine in the test framework
	docker run --rm -p 8080:8080 $(NAME)

#
#  Other targets
#

clean: ## Deletes all renewable artifacts, for build and install
	${DIR_RESOURCES}/clean.sh

compile: ## Compiles intermediate artifacts
	GOOS=linux GOARCH=amd64 go build -o ./bin/padelField

test-local: ## Run unit tests on local system (not container)
	go test -v ./...


