NAME := TestTask
ifeq ($(shell git tag --contains HEAD),)
  VERSION := $(shell git rev-parse --short HEAD)
else
  VERSION := $(shell git tag --contains HEAD)
endif
BUILDTIME := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
BUILDNAME := TestTask
GOLDFLAGS += -X TestTask/app/conf/buildvars.Version=$(VERSION)
GOLDFLAGS += -X TestTask/app/conf/buildvars.Buildtime=$(BUILDTIME)
GOLDFLAGS += -X TestTask/app/conf/buildvars.Buildname=$(BUILDNAME)
GOFLAGS = -ldflags "$(GOLDFLAGS)"

.PHONY: build

build: ## Build application
	GOSUMDB=off \
	go build -v -o TestTask $(GOFLAGS) ./cmd/server ; \

.DEFAULT_GOAL: build
