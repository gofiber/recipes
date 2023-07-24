.PHONY: all info go-build svelte-build install-template-dependencies clean

# Default target
all: svelte-build go-build

# Print information about available commands
info:
	$(info ------------------------------------------)
	$(info -           SvelteKit Embed App          -)
	$(info ------------------------------------------)
	$(info This Makefile helps you manage your projects.)
	$(info )
	$(info Available commands:)
	$(info - go-build:  Build the Golang project.)
	$(info - svelte-build:  Build the SvelteKit project.)
	$(info - all:  Run all commands (SvelteBuild, GoBuild).)
	$(info )
	$(info Usage: make <command>)

# Build the Golang project
go-build:
	@echo "=== Building Golang Project ==="
	@go build -o app -v

# Build the SvelteKit project
svelte-build: install-template-dependencies
	@echo "=== Building SvelteKit Project ==="
	@if command -v pnpm >/dev/null; then \
		pnpm run -C ./template build; \
	else \
		npm run --prefix ./template build; \
	fi

# Install template dependencies
install-template-dependencies:
	@if command -v pnpm >/dev/null; then \
		pnpm install -C ./template; \
	else \
		npm install --prefix ./template; \
	fi

# Clean build artifacts
clean:
	@echo "=== Cleaning build artifacts ==="
	@rm -f app
	@if [ -d "./template/__svelte_build__" ]; then \
		rm -rf ./template/__svelte_build__; \
	fi
