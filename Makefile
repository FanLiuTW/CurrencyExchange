SHELL := /bin/bash # Use bash syntax
.PHONY: doc install clean help
.DEFAULT: help

DOC_NAME=openapi.yaml

help:
	@echo "make install: compile packages and dependencies"
	@echo "make clean: remove object files and cached files"
	@echo "make doc: bundle a multi-file API definition into a single file"

doc:
	docker run --rm -v $(shell pwd):/mnt -w /mnt jeanberu/swagger-cli swagger-cli bundle doc/$(DOC_NAME) -r --outfile doc/spec.yaml --type yaml

install:
	@go build -v .
ifneq ($(findstring $(ENV),$(ENV_LIST)),)
	cp ./env/$(ENV).env .env
endif

clean:
	rm -f Asiayo
	go clean -i .