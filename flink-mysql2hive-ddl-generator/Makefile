.DEFAULT_GAOL := build

build: vet
	go build -o export
.PHONY:build

vet: fmt
	go vet ./...
.PHONY:vet

lint: fmt
	golint ./...
.PHONY:lint

fmt:
	go fmt ./...
.PHONY:fmt

