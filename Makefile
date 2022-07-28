.PHONY: default
default: help

.PHONY: help
help: ## help information about make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: up
up: ## provision Postgresql and run API using docker compose
	chmod +x ./db_scripts/01-init.sh
	docker-compose up

.PHONY: down
down: ## shutdown all services ran by docker compose
	docker-compose down

.PHONY: rm-volumes
rm-volumes: ## shutdown all services ran by docker compose and removes volumes
	docker-compose down --volumes

.PHONY: build
build: ## build services using docker compose
	docker-compose build
