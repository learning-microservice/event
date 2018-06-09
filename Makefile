# Makefile for event microservice

.DEFAULT_GOAL := build

# -----------------------------------------------------------------
#    ENV VARIABLE
# -----------------------------------------------------------------

NAME      := event
VERSION   := v1.0.0-alpha.1#$(shell git describe --tags --abbrev=0)
REVISION  := $(shell git rev-parse --short HEAD)

CMDDIR    := ./cmd
DESTDIR   := ./bin
SOURCEDIR := ./mvc
SOURCES   := $(shell find . -type f -name '*.go' | grep -v vendor)
LDFLAGS   := -ldflags="-s -w -X \"main.version=$(VERSION)\" -X \"main.revision=$(REVISION)\" -extldflags \"-static\""
NOVENDOR  := $(shell go list ./... | grep -v vendor)

BUILD_OPTS := -v -a -tags netgo -installsuffix netgo $(LDFLAGS) -o bin/$(NAME)

DOCKER_DB_NAME := $(NAME)-mysql-test

# -----------------------------------------------------------------
#    Main targets
# -----------------------------------------------------------------

bin/$(NAME): $(SOURCES)
	@go build $(BUILD_OPTS) $(CMDDIR)/api

clean: ## Remove temporary files
	@go clean
	@rm -rf bin/*

build: format bin/$(NAME) ## Build all libraries and binaries
	@true

format: ## Format all packages
	@goimports -w $(SOURCES)
	@go fmt $(NOVENDOR)

run: build ## Start the api server
	@bin/$(NAME)

test: docker-setup-testdb ## Run all the tests
	@export DB_ADDRESS="root:password@tcp(mysql:3306)/localdb?charset=utf8&parseTime=True&loc=Local"; go test -coverpkg=$(SOURCEDIR)/... -v $(NOVENDOR); make docker-teardown-testdb

lint: ## Code check
	@golint -set_exit_status $(CMDDIR)/...
	@golint -set_exit_status $(SOURCEDIR)/...
	@go vet $(NOVENDOR)

deps: ## Run dep ensure and prune
	@dep ensure -v
	@dep prune -v

update-deps: ## Run dep ensure update
	@dep ensure -update -v

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


# -----------------------------------------------------------------
#    Setup targets
# -----------------------------------------------------------------

setup: setup-dep setup-golint setup-goimports ## Setup dev tool (dep, golint and goimports)
	@true

setup-dep:
ifeq ($(shell command -v dep 2> /dev/null),)
	go get -u github.com/golang/dep/cmd/dep
endif

setup-golint:
ifeq ($(shell command -v golint 2> /dev/null),)
	go get -u github.com/golang/lint/golint
endif

setup-goimports:
ifeq ($(shell command -v goimports 2> /dev/null),)
	go get -u golang.org/x/tools/cmd/goimports
endif

# -----------------------------------------------------------------
#    Docker targets
# -----------------------------------------------------------------

docker-setup-testdb:
	@docker run -d --name $(DOCKER_DB_NAME) -e MYSQL_ROOT_PASSWORD=mysql -p "13306:13306" "mysql:5.7"

docker-teardown-testdb:
	@docker rm -f $(DOCKER_DB_NAME)

# -----------------------------------------------------------------
#    Release targets
# -----------------------------------------------------------------

install:
	go install $(LDFLAGS) $(SOURCEDIR)

release: ## Git tag and push version
	git tag $(VERSION)
	git push origin $(VERSION)

build-linux: format bin/$(NAME) ## Build all libraries and binaries for linux
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(BUILD_OPTS) $(CMDDIR)/api


.PHONY: $(/bin/bash egrep -o ^[a-zA-Z_-]+: $(MAKEFILE_LIST) | sed 's/://')