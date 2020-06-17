#!/bin/bash

ENV_OS=CGO_ENABLED=0 GOOS=linux GOARCH=amd64

.PHONY: demo_server
demo_server:
		$(ENV_OS) go build -o demo_server main.go

.PHONY: install_all
install_all: demo_server

.PHONY: docker_build
docker_build:install_all
	docker build -t go-server .