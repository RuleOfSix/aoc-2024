.DEFAULT_GOAL := build
.PHONY: fmt vet build clean

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -gcflags "-N" p1/part_one.go 

clean:
	@rm part_one part_two 2>/dev/null || true
