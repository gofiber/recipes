.PHONY: build

# What CI runs, see .github/CONTRIBUTING.md
build:
	npm ci --prefix web
	npm test --prefix web
	npm run build --prefix web
	go build ./...
