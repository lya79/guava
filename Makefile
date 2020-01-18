PATH_OF_DOCKER_COMPOSE_YML=./deployment/local/docker-compose.yml
GOMOD=go.mod

help:
	@echo "指令說明:"
	@echo "make: 檢視目前命令說明."
	@echo "make install: 開啟 go module功能並且安裝第三方套件."
	@echo "make test: 使用 go run執行全部服務."
	@echo "make up: 使用 docker-compose啟動全部服務."
	@echo "make down: 使用 docker-compose停止全部服務."
	@echo "make build: 使用 docker-compose build全部服務."
	@echo "make run: 使用 go run執行全部服務."

install:
	export GO111MODULE=on
ifeq ($(GOMOD), $(wildcard $(GOMOD)))
	@echo "go.mod已經存在."
else
	go mod init
endif
	go mod tidy 
	go mod vendor

test: install
	go vet ./...
	docker-compose -f $(PATH_OF_DOCKER_COMPOSE_YML) up -d mysql redis 
	docker-compose -f $(PATH_OF_DOCKER_COMPOSE_YML) up -d adminer redis-admin 
	docker-compose -f $(PATH_OF_DOCKER_COMPOSE_YML) ps
	PROJECT_ROOT=$(shell pwd) PROJECT_ENV=local PROJECT_SITE=chatroom go test -v ./...
	docker-compose -f $(PATH_OF_DOCKER_COMPOSE_YML) down 

up: install test
	docker-compose -f $(PATH_OF_DOCKER_COMPOSE_YML) up -d
	docker-compose -f $(PATH_OF_DOCKER_COMPOSE_YML) ps

down: install 
	docker-compose -f $(PATH_OF_DOCKER_COMPOSE_YML) down 
	docker-compose -f $(PATH_OF_DOCKER_COMPOSE_YML) ps

build: install test 
	docker-compose -f $(PATH_OF_DOCKER_COMPOSE_YML) up -d --build
	docker-compose -f $(PATH_OF_DOCKER_COMPOSE_YML) ps

run: install test 
	docker-compose -f $(PATH_OF_DOCKER_COMPOSE_YML) up -d mysql redis 
	docker-compose -f $(PATH_OF_DOCKER_COMPOSE_YML) up -d adminer redis-admin 
	docker-compose -f $(PATH_OF_DOCKER_COMPOSE_YML) ps
	PROJECT_ROOT=$(shell pwd) PROJECT_ENV=local PROJECT_SITE=chatroom go run .