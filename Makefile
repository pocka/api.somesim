.SUFFIXES:

.SUFFIXES: .go


# Main rules
.PHONY: default
default: all

.PHONY: all
all:


# Docker rules
DOCKER_LABEL_DEVELOP = pocka/api.somesim-dev
DOCKER_LABEL_SERVER = pocka/api.somesim
DOCKER_RUN = docker run -v $(shell pwd):/work -u $(shell id -u):$(shell id -g)

## Container creation
.PHONY: container
container: container/develop container/server

.PHONY: container/develop
container/develop: docker/develop/Dockerfile
	docker build -t $(DOCKER_LABEL_DEVELOP) -f $< .

.PHONY: container/server
container/server: docker/server/Dockerfile
	$(DOCKER_RUN) --entrypoint 'make' $(DOCKER_LABEL_DEVELOP) dist
	docker build -t $(DOCKER_LABEL_SERVER) -f $< .

## Run container
.PHONY: run/develop
run/develop:
	-$(DOCKER_RUN) -it $(DOCKER_LABEL_DEVELOP)

.PHONY: run/server
run/server:
	-docker run $(DOCKER_LABEL_SERVER)

# Compile rules
# !! DO NOT RUN THESE RULES IN THE OUTSIDE OF CONTAINER !!

dist: dist/somesim

GO_FILES = $(shell find src -name "*.go")
GO_BUILD = CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo

## Build binary
dist/somesim: $(GO_FILES)
	$(GO_BUILD) -o $@ src/main.go
	chmod +x $@
