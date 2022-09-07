# Name of this Docker-Container
NAME = padel-field

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

build: clean # Build the docker image
	 docker build . -t $(NAME) -f 'Dockerfile'

run: # start the app
	docker run --rm -p 8080:8080 $(NAME)

start-db:
	docker run -e POSTGRES_USER=docker -e POSTGRES_PASSWORD=root -e POSTGRES_DB=padelfield library/postgres



#
#  Other targets
#

clean: ## Deletes all renewable artifacts, for build and install
	docker rmi -f $(NAME):latest || echo "No such image:" \
	.$(DIR_RESOURCES)/clean.sh


test-local: ## Run unit tests on local system (not container)
	go test -v ./...


