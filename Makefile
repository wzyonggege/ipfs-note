all: fmt vet build

build: build-ipfs

build-ipfs:
	@echo "==== go build ==="
	@go build -o ipfsapp -v main.go

build-ipfs-linux:
	@echo "go build linux version"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ipfsapp -v main.go

clean:
	rm -f ipfsapp
	find . -name "[._]*.s[a-w][a-z]" | xargs -i rm -f {}

vet:
	@echo "==== go vet ==="
	@go vet $(go list ./... | grep -v /vendor/)

fmt:
	@echo "=== go fmt ==="
	@go fmt $(go list ./... | grep -v /vendor/)

.PHONY: clean fmt vet build
