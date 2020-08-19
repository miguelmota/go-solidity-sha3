all:
	@echo "no default"

.PHONY: test
test:
	@go test -v `go list ./... | grep -v example` $(ARGS) && echo "ALL PASS" || echo "FAILURE"

.PHONY: deps
deps:
	@GO111MODULE=on go mod vendor

.PHONY: release
release:
	@rm -rf dist
	@goreleaser
