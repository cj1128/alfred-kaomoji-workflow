SHELL := /bin/bash

install:
	go install
.PHONY: install

build:
	go build -o workflow/kaomoji
.PHONY: build

bundle: build
	upx --brute workflow/kaomoji
	cd workflow && zip -r ../tmp/Kaomoji.alfredworkflow . && cd ..
.PHONY: bundle
