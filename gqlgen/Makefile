.PHONY: generate build

# Regenerates graph/generated.go, graph/model and resolver stubs from graph/schema.graphqls
generate:
	go generate ./...

# CI runs this target (see .github/CONTRIBUTING.md). It fails if go generate changes anything in graph/.
build:
	@before="$$(cat graph/*.go graph/model/*.go | cksum)"; \
	go generate ./... || exit 1; \
	if [ "$$before" != "$$(cat graph/*.go graph/model/*.go | cksum)" ]; then \
		echo "generated code was out of date: commit the regenerated files in graph/" >&2; exit 1; \
	fi
	go vet ./...
	go build ./...
