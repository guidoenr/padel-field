# API
API_CONTAINER_NAME = go-gin-api
API_IMAGE_NAME = go-gin-api-image
API_IP = 172.20.0.4
# DB
DB_CONTAINER_NAME = db-postgres
DB_IMAGE_NAME = postgres:10-alpine
DB_IP = 172.20.0.5

# NETWORK
NETWORK_NAME = network-gin-postgres
NETWORK_SUBNET = 172.20.0.0/16

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


start-env: ## start the network
	docker network rm $(NETWORK_NAME)
	docker network create --subnet $(NETWORK_SUBNET) $(NETWORK_NAME)


start-api: build-api ## start the api
	docker run --rm --name $(API_CONTAINER_NAME) --network=$(NETWORK_NAME) --ip=$(API_IP) -p 8080:8080 $(API_IMAGE_NAME)


start-db: ## start the db
	docker run --rm --name $(DB_CONTAINER_NAME) --network=$(NETWORK_NAME) --ip=$(DB_IP) -P -p 6543:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -e POSTGRES_DB=padelfield $(DB_IMAGE_NAME)


build-api: ## build the api docker image
	docker build . -t $(API_IMAGE_NAME) -f 'Dockerfile'

#
#  Other targets
#

clean-images: ## deletes all renewable artifacts, for build and install
	docker rmi -f $(API_IMAGE_NAME):latest
	docker rmi -f $(DB_IMAGE_NAME):latest

test-local: ## run unit tests on local system (not container) - no tests at the moment
	go test -v ./...


