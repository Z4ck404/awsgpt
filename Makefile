NAME 	= $(shell go list -m)

build: clean
	go build -o bin/awsgpt

clean:
	rm -rf bin/awsgpt*

fmt: ## Run go fmt against code.
	$(GOIMPORTS) -w -l -local $(NAME) $(shell pwd)
	go fmt ./...


lint: ## Run golangci-lint against code.
	$(GOLANGCI_LINT) run --max-same-issues 0  --timeout=5m0s $(shell pwd)/...


cross-build: OSARCH = linux/amd64 darwin/amd64 linux/arm64 darwin/arm64
cross-build: clean
	${GOX} -output="bin/awsgpt_{{.OS}}_{{.Arch}}" -osarch="${OSARCH}"

GOX = CGO_ENABLED=0 go run github.com/mitchellh/gox@latest
GOIMPORTS = go run golang.org/x/tools/cmd/goimports@v0.12.0
GOLANGCI_LINT = go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3
