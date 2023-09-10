# Database Configuration
DB_NAME ?= postgres
DB_USER ?= postgres
DB_PASSWORD ?= mysecretpassword
DB_HOST ?= localhost
DB_PORT ?= 5432
DB_TYPE ?= postgres

# PSQL URL
PSQLURL ?= $(DB_TYPE)://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)

# Container Configuration
NET_NAME ?= go-postgres
CON_NAME ?= postgres_db
POST_VERSION ?= postgres:15.2-alpine

.PHONY: postgresup postgresdown psql test build go_app run delete_container delete_image

postgresup:
	docker run \
	--name $(CON_NAME) \
	--network $(NET_NAME) \
	-p 5433:5432 \
	-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
	-v $(PWD):/var/lib/postgresql/data \
	-d $(POST_VERSION)

postgresdown:
	docker stop $(CON_NAME) && docker rm $(CON_NAME)

psql:
	docker exec -it $(CON_NAME) psql $(PSQLURL)

test:
	go test ./test -v

build:
	docker build -t go-rest-api:0.0.1 .

go_app:
	docker run --name go-rest-api \
	-p 8181:8181 \
	--network $(NET_NAME) \
	-d -t go-rest-api:0.0.1

run: go_app postgresup

delete_container:
	docker rm -f $(CON_NAME)

delete_image:
	docker rmi -f go-rest-api:0.0.1
