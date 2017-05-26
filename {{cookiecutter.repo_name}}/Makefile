SHELL := /bin/bash
PWD = $(shell pwd)
GREP := $(shell command -v ggrep || command -v grep)

help:
	@$(GREP) --only-matching --word-regexp '^[^[:space:].]*:' Makefile | sed 's|:[:space:]*||'

rpi:
	GOOS=linux GOARCH=arm GOARM=7 go build -v .

clean:
	-rm {{ cookiecutter.repo_name }}

.PHONY: clean help
