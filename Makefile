VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
HTTPS_GIT := https://github.com/JackalLabs/jlaunch.git

ldflags = -X github.com/JackalLabs/jlaunch/cmd.Version=$(VERSION) \
		  -X github.com/JackalLabs/jlaunch/cmd.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)' -trimpath


install: tidy
	@go install $(BUILD_FLAGS) ./
	@jlaunch version

build: tidy
	@go build -o build/jlaunch -ldflags '$(LD_FLAGS)' ./

tidy:
	@go mod tidy


format-tools:
	go install mvdan.cc/gofumpt@v0.5.0
	gofumpt -l -w .


lint: format-tools
	golangci-lint run

format: format-tools
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/lcd/statik/statik.go" | xargs gofumpt -w -s
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/lcd/statik/statik.go" | xargs misspell -w
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/lcd/statik/statik.go" | xargs goimports -w -local github.com/jackalLabs/canine-chain



.PHONY: install format lint format-tools tidy clean build