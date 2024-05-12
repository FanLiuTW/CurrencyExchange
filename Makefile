SHELL := /bin/bash # Use bash syntax
.PHONY: install clean help
.DEFAULT: help

help:
	@echo "make install: compile packages and dependencies"
	@echo "make clean: remove object files and cached files"

install:
	@go build -v .
ifneq ($(findstring $(ENV),$(ENV_LIST)),)
	cp ./env/$(ENV).env .env
endif

clean:
	rm -f Asiayo
	go clean -i .