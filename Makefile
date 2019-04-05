APPLICATION_NAME := $(shell grep "const ApplicationName " version.go | sed -E 's/.*"(.+)"$$/\1/')
BIN_NAME=${APPLICATION_NAME}

VERSION := $(shell grep "const Version " version.go | sed -E 's/.*"(.+)"$$/\1/')
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
default: test

help:
	@echo 'Management commands for ${APPLICATION_NAME}:'
	@echo
	@echo 'Usage:'
	@echo '    make get-deps            Runs glide install'
	@echo '    make up-deps             Runs glide update'
	@echo '    make run-test                Run tests on a compiled project.'
	@echo

get-deps:
	glide install

up-deps:
	glide up

run-test:
	mkdir -p ./test/cover
	go test -race -coverpkg= ./... -coverprofile=./test/cover/cover.out
	go tool cover -html=./test/cover/cover.out -o ./test/cover/cover.html
