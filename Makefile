.DEFAULT_GOAL := dev

.PHONY: dev
dev: ## dev build
dev: clean install generate buildweb vet fmt spell lint test mod-tidy

.PHONY: ci
ci: ## CI build
ci: dev diff

.PHONY: clean
clean: ## remove files created during build pipeline
	$(call print-target)
	rm -rf dist
	rm -f coverage.*

.PHONY: install
install: ## go install tools
	$(call print-target)
	go install tool

.PHONY: generate
generate: ## go generate
	$(call print-target)
	go generate ./...

.PHONY: buildweb
buildweb: ## generate webapp build artifacts
	$(call print-target)
	cd web/app && yarn && yarn build

.PHONY: vet
vet: ## go vet
	$(call print-target)
	go vet ./...

.PHONY: fmt
fmt: ## go fmt
	$(call print-target)
	go fmt ./...
	cd web/app && yarn format

.PHONY: spell
spell: ##misspell
	$(call print-target)
	go tool misspell -error -locale=US -w **.md

.PHONY: lint
lint: ## golangci-lint
	$(call print-target)
	go tool golangci-lint run
	cd web/app && yarn lint

.PHONY: test
test: ## go test with race detector and code covarage
	$(call print-target)
	go test -race -covermode=atomic -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

.PHONY: mod-tidy
mod-tidy: ## go mod tidy
	$(call print-target)
	go mod tidy

.PHONY: diff
diff: ## git diff
	$(call print-target)
	git diff --exit-code
	RES=$$(git status --porcelain) ; if [ -n "$$RES" ]; then echo $$RES && exit 1 ; fi

.PHONY: build
build: ## goreleaser --snapshot --skip-publish --clean
build: install buildweb
	$(call print-target)
	go tool goreleaser --snapshot --skip-publish --clean

.PHONY: release
release: ## goreleaser --clean
release: install buildweb
	$(call print-target)
	go tool goreleaser --clean

.PHONY: run
run: ## go run
	@go run -race ./cmd/deviceportal

.PHONY: go-clean
go-clean: ## go clean build, test and modules caches
	$(call print-target)
	go clean -r -i -cache -testcache -modcache

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

define print-target
    @printf "Executing target: \033[36m$@\033[0m\n"
endef
