OK_COLOR=\033[32;01m
NO_COLOR=\033[0m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

GO := go
GO_MOD := go mod

ROOT_DIR := $(realpath .)

PKGS = $(shell $(GO) list ./...)
GOYAPH = ${GOPATH}

BIN_WEBSERVER := api

build:
	@echo "$(OK_COLOR)==> Building binary (/$(BIN_WEBSERVER))...$(NO_COLOR)"
	@echo $(GO) build -v -o /$(BIN_WEBSERVER) ./cmd/
	@$(GO) build -v -o $(BIN_WEBSERVER) ./cmd/

develop:
	@docker-compose up