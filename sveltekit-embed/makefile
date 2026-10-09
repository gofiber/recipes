.PHONY: all build info go-build svelte-build install-template-dependencies clean

# pinned so the lockfile format stays stable; npx needs nothing but Node
PNPM ?= npx --yes pnpm@10.34.6

# Default target
all: svelte-build go-build

# What CI runs, see .github/CONTRIBUTING.md
build: all

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
	$(info - build:  Same as all, the target CI runs.)
	$(info )
	$(info Usage: make <command>)

# Build the Golang project
go-build:
	@echo "=== Building Golang Project ==="
	@go build -o app -v

# Build the SvelteKit project
svelte-build: install-template-dependencies
	@echo "=== Building SvelteKit Project ==="
	@$(PNPM) run -C ./template build

# Install template dependencies
install-template-dependencies:
	@$(PNPM) install -C ./template --frozen-lockfile

# Clean build artifacts
clean:
	@echo "=== Cleaning build artifacts ==="
	@rm -f app
	@find ./template/dist -mindepth 1 ! -name .gitkeep -delete
