TAG := $(shell git describe --tags)

ROOT=$(shell pwd)
CMD= $(ROOT)/cmd
WEB= $(CMD)/web
WORKER= $(CMD)/worker
MIGRATE= $(CMD)/migrate

web-build:
	ROOT=$(ROOT) $(MAKE) -C $(WEB)

web-dev:
	ROOT=$(ROOT) air -c cmd/web/air.toml

worker-build:
	ROOT=$(ROOT) $(MAKE) -C $(WORKER)

worker-dev:
	ROOT=$(ROOT) $(MAKE) -C $(WORKER) dev

migrate-build:
	ROOT=$(ROOT) $(MAKE) -C $(MIGRATE) build

migrate-create:
	migrate create -ext sql -dir db/migrations $(name)

unit-test:
	go test ./... -v -cover
