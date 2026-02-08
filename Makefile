## help: 💡 Display available commands
.PHONY: help
help:
	@echo '⚡️ GoFiber/Recipes Development:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## audit: 🚀 Conduct quality checks
.PHONY: audit
audit:
	go mod verify
	go vet ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

## format: 🎨 Fix code format issues
.PHONY: format
format:
	go run mvdan.cc/gofumpt@latest -w -l .

## markdown: 🎨 Find markdown format issues (Requires markdownlint-cli)
.PHONY: markdown
markdown:
	markdownlint-cli2 "**/*.md" "#vendor"

## lint: 🚨 Run lint checks
.PHONY: lint
lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0 run ./...

## generate: ⚡️ Generate implementations
.PHONY: generate
generate:
	go generate ./...
