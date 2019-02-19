# Golang staff
GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean -cache
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install

# PKG
PKG_NAME = protoc-go-inject-tag

# Environment
#include .env
#export

# PATH's
SRC=$(abspath .)

install:
	$(GOGET) ./...
	$(GOBUILD) -v -i -ldflags "-X main.GitCommit=`(git rev-list -1 HEAD)`" -o $(GOPATH)/bin/$(PKG_NAME)

generate:
	$(clean)
	cd $(SRC)/pb && protoc -I . --go_out=. *.proto && \
		$(PKG_NAME) -I .

clean:
	rm -rf $(SRC)/pb/*.go
