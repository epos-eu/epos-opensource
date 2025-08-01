BIN=epos-opensource
VERSION?=makefile
GOOS?=$(shell go env GOOS)
GOARCH?=$(shell go env GOARCH)

build: generate
	go build -ldflags "-X github.com/epos-eu/epos-opensource/cmd.Version=$(VERSION)" -o $(BIN) .

# Build for specific platform (used by CI)
build-release:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-X github.com/epos-eu/epos-opensource/cmd.Version=$(VERSION)" -o $(BIN)-$(GOOS)-$(GOARCH)$(EXT) .

clean:
	rm -f $(BIN)*

generate:
	go generate ./...

lint:
	golangci-lint run ./...

test:
	go test ./...

.PHONY: build build-release clean generate lint test
