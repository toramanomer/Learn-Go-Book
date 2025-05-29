# Each possible operation is called a target.
# The .DEFAULT_GOAL defines which target is run when no target is specified.
# You must indent the steps in a target with a tab.
.DEFAULT_GOAL := build

.PHONY: fmt vet build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build

clean:
	go clean
