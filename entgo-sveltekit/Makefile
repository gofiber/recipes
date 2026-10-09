.PHONY: build

# pinned so the lockfile format stays stable; npx needs nothing but Node
PNPM ?= npx --yes pnpm@10.34.6

# What CI runs, see .github/CONTRIBUTING.md
build:
	$(PNPM) install -C ./template --frozen-lockfile
	$(PNPM) run -C ./template build
	go build -o app -v
