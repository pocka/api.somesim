.SUFFIXES:

.SUFFIXES: .go


# Main rules
.PHONY: default
default: all

.PHONY: all
all: container/server

.PHONY: init
init: .crt
	mkdir -p dist

CRT_DIR = $(shell pwd)/.crt
CRT_NAME = testserver
.crt:
	mkdir -p .crt
	cd /etc/pki/tls/certs && \
	make $(CRT_DIR)/$(CRT_NAME).crt
	openssl rsa -in $(CRT_DIR)/$(CRT_NAME).key -out $(CRT_DIR)/$(CRT_NAME).key
	chmod +r .crt/*

.PHONY: clean
clean:
	rm -rf .crt
	rm -rf dist

# Docker rules
DOCKER_LABEL_DEVELOP = pocka/api.somesim-dev
DOCKER_LABEL_SERVER = pocka/api.somesim
DOCKER_LABEL_TESTSERVER = pocka/api.somesim-test
DOCKER_RUN = docker run --rm -v $(shell pwd)/src:/go/src/github.com/pocka/api.somesim -v $(shell pwd):/work -u $(shell id -u):$(shell id -g)

## Container creation
.PHONY: container/develop
container/develop: docker/develop/Dockerfile
	docker build -t $(DOCKER_LABEL_DEVELOP) -f $< .

.PHONY: container/server
container/server: docker/server/Dockerfile dist docs/index.html
	docker build -t $(DOCKER_LABEL_SERVER):v1 -f $< .

## Run container
.PHONY: run/develop
run/develop:
	-$(DOCKER_RUN) --entrypoint '/bin/sh' -it $(DOCKER_LABEL_DEVELOP)

.PHONY: run/server
run/server:
	-docker run --rm -v $(CRT_DIR):/work/.crt -p 8080:80 $(DOCKER_LABEL_SERVER):v1 --tls-certificate=/work/.crt/$(CRT_NAME).crt --tls-key=/work/.crt/$(CRT_NAME).key

# Compile rules

dist: dist/somesim

GO_FILES = $(shell find src -name "*.go")
GO_BUILD = go build -a -installsuffix netgo -tags netgo --ldflags '-extldflags "-static"'

## Build binary
dist/somesim: $(GO_FILES)
	$(DOCKER_RUN) -e "CGO_ENABLED=0" -e "GOOS=linux" -w '/go/src/github.com/pocka/api.somesim' $(DOCKER_LABEL_DEVELOP) $(GO_BUILD) -o /work/$@ cmd/somesim-server/main.go
	chmod +x $@

# Swagger stuff
.PHONY: swagger/validate
swagger/validate: docs/swagger.yml
	$(DOCKER_RUN) $(DOCKER_LABEL_DEVELOP) swagger validate $<

.PHONY: swagger/generate
swagger/generate: docs/swagger.yml
	$(DOCKER_RUN) -w '/go/src/github.com/pocka/api.somesim' $(DOCKER_LABEL_DEVELOP) swagger generate server -f /work/$< -A somesim --exclude-main -t ./
