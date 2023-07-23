GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test -v --cover
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run
GOMODTIDY=$(GOCMD) mod tidy
GOVET=$(GOCMD) vet
GOIMPORTS=goimports

.PHONY: run
.SILENT:
