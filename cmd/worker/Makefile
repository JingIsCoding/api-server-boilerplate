VERSION := 0.0.1
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME := worker

build:
	go build -o ../../bin

dev:
	ENV=local ROOT=$(ROOT) air -c air.toml
