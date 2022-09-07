# API
API_CONTAINER_NAME = go-gin-api
API_IMAGE_NAME = go-gin-api-image

# DB
DB_CONTAINER_NAME = db-postgres
DB_IMAGE_NAME = postgres:alpine

# NETWORK
NETWORK_NAME = network-gin-postgres

# Relative path to bin (scripts) directory (you may need to add/remove ".." paths)
DIR_API = /api
DIR_UI = /ui
DIR_DB = /models/psdb/
DIR_RESOURCES = /resources

.DEFAULT_GOAL := help
help: ## Print the list of makefile targets
	@grep -hE '^[%a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "%-18s %s\n", $$1, $$2}'

#
#  BUILD AND RUN
#

run: build-api start-api start-db create-network

build-api: # Build the docker image
	docker build . -t $(API_IMAGE_NAME) -f 'Dockerfile'


start-api: build-api # start the api
	docker run --rm --name $(API_CONTAINER_NAME) -p 8080:8080 $(API_IMAGE_NAME)


start-db: # start the db
	docker run --rm --name $(DB_CONTAINER_NAME) -P -p 127.0.0.1:5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -e POSTGRES_DB=padelfield $(DB_IMAGE_NAME)


create-network: # start the network and connect the containers
	docker network rm $(NETWORK_NAME) || echo "No network"
	docker network create $(NETWORK_NAME)
	docker network connect $(NETWORK_NAME) $(API_CONTAINER_NAME)
	docker network connect $(NETWORK_NAME) $(DB_CONTAINER_NAME)

#
#  Other targets
#


clean-images: ## Deletes all renewable artifacts, for build and install
	docker rmi -f $(API_IMAGE_NAME):latest
	docker rmi -f $(DB_IMAGE_NAME):latest

test-local: ## Run unit tests on local system (not container)
	go test -v ./...


