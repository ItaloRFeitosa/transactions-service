LOCAL_COMPOSE_FILE := ./deployments/docker-compose.local.yml
LOCAL_PROJECT_NAME := transactions-service-local

COMPOSE_FILE := ./deployments/docker-compose.yml
PROJECT_NAME := transactions-service

.PHONY: build up stop logs up_build destroy dev

build:
	go build -o ./out/api ./cmd/api/main.go

up:
	DOCKER_BUILDKIT=0 docker-compose --project-directory . -p $(compose_project_name) -f $(compose_file) up -d

stop:
	DOCKER_BUILDKIT=0 docker-compose --project-directory . -p $(compose_project_name) -f $(compose_file) stop

logs:
	DOCKER_BUILDKIT=0 docker-compose --project-directory . -p $(compose_project_name) -f $(compose_file) logs api --follow

up_build:
	DOCKER_BUILDKIT=0 docker-compose --project-directory . -p $(compose_project_name) -f $(compose_file) up -d --build

destroy:
	DOCKER_BUILDKIT=0 docker-compose --project-directory . -p $(compose_project_name) -f $(compose_file) down -v

local:
	$(MAKE) up compose_project_name=${LOCAL_PROJECT_NAME} compose_file=$(LOCAL_COMPOSE_FILE)

local_build:
	$(MAKE) up_build compose_project_name=${LOCAL_PROJECT_NAME} compose_file=$(LOCAL_COMPOSE_FILE)

local_logs:
	$(MAKE) logs compose_project_name=${LOCAL_PROJECT_NAME} compose_file=$(LOCAL_COMPOSE_FILE)

local_destroy:
	$(MAKE) destroy compose_project_name=${LOCAL_PROJECT_NAME} compose_file=$(LOCAL_COMPOSE_FILE)

prodlike:
	$(MAKE) up_build compose_project_name=${PROJECT_NAME} compose_file=$(COMPOSE_FILE)

prodlike_logs:
	$(MAKE) logs compose_project_name=${PROJECT_NAME} compose_file=$(COMPOSE_FILE)

prodlike_destroy:
	$(MAKE) destroy compose_project_name=${PROJECT_NAME} compose_file=$(COMPOSE_FILE)