# Use webrpc-gen binary from the root of the project.
export PATH = $(shell echo $$PWD/../../bin:$$PATH)

# The virtualenv from `make venv`, if present, otherwise python3 on the PATH.
PYTHON ?= $(if $(wildcard .venv/bin/python),.venv/bin/python,python3)

.PHONY: all build generate venv test

all:
	@awk -F'[ :]' '!/^all:/ && /^([A-z_-]+):/ {print "make " $$1}' Makefile

build:
	make -C ../../ build

generate: build
	webrpc-gen -schema=./example.ridl -target=python -client -server -out=./example.gen.py

# Create ./.venv with the generated client's dependency.
venv:
	python3 -m venv .venv
	.venv/bin/pip install --quiet httpx

# End-to-end tests of the Python generator (requires Python 3.9+ and httpx).
test: build
	PYTHON=$(PYTHON) ./test.sh
