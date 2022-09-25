# Is defined if `swag` command (used to update API docs) exists
SWAGGER_EXISTS := $(shell which swag 2>/dev/null)
# A warning message that will be printed ifndef SWAGGER_EXISTS
define SWAGGER_ERROR
swag command is missing.

The swagger documentation won't be updated, which may lead to outdated API docs.
To install swag, run:
  make swag-install

Refer to https://github.com/swaggo/swag#getting-started for details.
If you've already installed swag, check if it can be accessed in your PATH environment.
endef


.PHONY: run
run:
	go run ./cmd/main.go

.PHONY: swag-install
swag-install:
	go get -u github.com/swaggo/swag/cmd/swag
	go install github.com/swaggo/swag/cmd/swag

.PHONY: swag
swag:
ifdef SWAGGER_EXISTS
	swag init -g ./cmd/main.go
else
	$(warning $(SWAGGER_ERROR))
endif
