TAG := $(shell git describe --tags)

BASE=$(shell pwd)
CMD= $(BASE)/cmd
WEB= $(CMD)/web

web:
	$(MAKE) $(ARGS) -C $(WEB)
