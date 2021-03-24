NPROCS = $(shell grep -c 'processor' /proc/cpuinfo)
MAKEFLAGS += -j$(NPROCS)

.PHONY: clean
clean:
	@rm -fv ./testoauth2
	@rm -fv ./logs/*


.PHONY: run
run: clean
	@GOOS=linux go build -o ./testoauth2 ./*.go
	./testoauth2

.PHONY: build
build: clean
	@GOOS=linux go build -o ./testoauth2 ./*.go

.PHONY: buildrelease
buildrelease:
	@GOOS=linux go build -ldflags="-s -w" -o ./testoauth2 ./*.go
	@upx --brute ./testoauth2

.PHONY: release
release: clean buildrelease