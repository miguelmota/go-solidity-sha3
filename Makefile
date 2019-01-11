all:
	@echo "no default"

.PHONY: test
test:
	go test -v ./...

.PHONY: deps
deps:
	@GO111MODULE=1 go mod vendor
