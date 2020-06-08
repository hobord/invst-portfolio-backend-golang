-include .env

# VERSION := $(shell git describe --tags)
# BUILD := $(shell git rev-parse --short HEAD)
TARGET := bin/portfolio-server


# Use linker flags to provide version/build settings
# LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

.DEFAULT_GOAL := default
configure:
	go mod tidy
	go get github.com/vektra/mockery/.../
	mockery -all
test:
	go test -race 

compile:
	go build -o $(TARGET)

default: compile
