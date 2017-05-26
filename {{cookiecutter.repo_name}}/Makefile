SHELL := /bin/bash
PWD = $(shell pwd)
GREP := $(shell command -v ggrep || command -v grep)

help:
	@$(GREP) --only-matching --word-regexp '^[^[:space:].]*:' Makefile | sed 's|:[:space:]*||'

clean:
	-rm {{ cookiecutter.repo_name }}

.PHONY: clean help
