SHELL := /bin/bash
PWD = $(shell pwd)
GREP := $(shell command -v ggrep || command -v grep)
GIT := $(shell command -v /usr/local/bin/git || command -v git)

LDFLAGS := -ldflags "-X main.version=$$($(GIT) describe --tags)"

help:
	@$(GREP) --only-matching --word-regexp '^[^[:space:].]*:' Makefile | sed 's|:[:space:]*||'

build:
	go build $(LDFLAGS) -v .

test:
	go test -v .

bench:
	go test -bench=.

rpi:
	GOOS=linux GOARCH=arm GOARM=7 go build $(LDFLAGS) -v .

oldrpi:
	GOOS=linux GOARCH=arm GOARM=6 go build $(LDFLAGS) -v .

macos:
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -v .

linux:
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -v .

clean:
	-rm {{ cookiecutter.repo_name }}

.PHONY: clean help
