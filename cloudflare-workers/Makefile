.PHONY: build

# pinned for reproducible installs; npx needs nothing but Node
BUN ?= npx --yes bun@1.4.2

# What CI runs, see .github/CONTRIBUTING.md
build:
	$(BUN) install --frozen-lockfile
	$(BUN) x tsc --noEmit
