SHELL := /bin/bash
PWD = $(shell pwd)
GREP := $(shell command -v ggrep || command -v grep)

.PHONY: clean help

help:
	@$(GREP) --only-matching --word-regexp '^[^[:space:].]*:' Makefile | sed 's|:[:space:]*||'

clean:
	@echo
