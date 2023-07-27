cnf ?= config.env
include $(cnf)
export $(shell sed 's/=.*//' $(cnf))
# DOCKER TASKS
# Build the container
build: ## Build the container
	docker build -t carsim .

build-nc: ## Build the container without caching
	docker build --no-cache -t $(APP_NAME) .

run: ## Run container on port configured in `config.env`
	docker run -i -t --rm --env-file=./config.env --name="$(APP_NAME)" $(APP_NAME)

test:
	go test wheels/wheels_test.go