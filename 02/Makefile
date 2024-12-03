.DEFAULT_GOAL := build
.PHONY: fmt vet build clean

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -gcflags "-N" p1/part_one.go 
	go build -gcflags "-N" p2/part_two.go

clean:
	@go clean
