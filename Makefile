#include ./src/.env

all: up

init:
	-mkdir ./libs
	GOPATH=$$(echo "$${PWD}/libs") go build ./src/main.go

up: build
	docker compose -f ./deploy/docker/docker-compose.yml -p pspbalsaas up -d

build:
	docker build -t pspbalsaas-image .

down:
	docker compose -f ./deploy/docker/docker-compose.yml -p pspbalsaas down

deploy:
	ansible-playbook ./deploy/docker/playbook-up.yaml

destroy:
	ansible-playbook ./deploy/docker/playbook-down.yaml

db-term:
	docker exec -it pspbalsaas-db-1 bash

redis-cli:
	docker exec -it pspbalsaas-db-1 redis-cli

restart: down
	docker compose -f ./deploy/docker/docker-compose.yml -p pspbalsaas up -d

