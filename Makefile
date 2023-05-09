.PHONY: help build run clean
.DEFAULT_GOAL := help
.SILENT: help build run clean

build:
	@echo "Building binary..."
	@go build -o bin/localk8s -v

run: build
	@echo "Running binary..."
	@./bin/localk8s -action $(ACTION)

clean:
	@echo "Cleaning..."
	@rm -f $(shell ls -a ./bin/*  | grep -v '\.gitkeep')

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build		Build binary, use with BINARY_NAME=xxx"
	@echo "  run		Run binary, use with ACTION=create/delete"
	@echo "  clean		Clean binary"
	@echo "  help		Show this help"
