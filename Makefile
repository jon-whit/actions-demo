######################################################
#
# Development targets
#
######################################################

.PHONY: download
download:
	@go mod download

.PHONY: install-tools
install-tools: download
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

.PHONY: generate
generate: buf-generate  go-generate

.PHONY: go-generate
go-generate: install-tools
	@go generate ./...

.PHONY: buf-generate
buf-generate: install-tools
	@buf generate

.PHONY: test
test: generate
	@go test -v -race ./...

.PHONY: test-short
test-short: generate
	@go test -v -race -short ./...

.PHONY: test-coverage
test-coverage: generate
	@go test -v -race -coverprofile=coverage.out ./...

.PHONY: check
check: check-vet check-lint

.PHONY: check-vet
check-vet:
	@go vet ./...

.PHONY: check-lint
check-lint:
	@staticcheck ./...

ci-test-coverage:
	@go test -race -coverprofile=coverage.txt -covermode=atomic ./...