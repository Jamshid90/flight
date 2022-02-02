include .env
export

CURRENT_DIR=$(shell pwd)

APP=flight

CMD_DIR=./cmd

.DEFAULT_GOAL = build

# build for current os
.PHONY: build
build:
	go build -ldflags="-s -w" -o ./bin/${APP} ${CMD_DIR}/main.go

# build for linux amd64
.PHONY: build-linux
build-linux:
	CGO_ENABLED=0 GOARCH="amd64" GOOS=linux go build -ldflags="-s -w" -o ./bin/${APP} ${CMD_DIR}/main.go

# run service
.PHONY: run
run:
	go run ${CMD_DIR}/main.go

# migrate	
.PHONY: migrate
migrate:
	migrate -source file://migrations -database postgresql://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=disable up
