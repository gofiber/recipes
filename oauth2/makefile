NPROCS = $(shell grep -c 'processor' /proc/cpuinfo)
MAKEFLAGS += -j$(NPROCS)

.PHONY: clean
clean:
	@rm -fv ./oauth2
	@rm -fv ./logs/*


.PHONY: run
run: clean
	@GOOS=linux go build -o ./oauth2 ./*.go
	./oauth2

.PHONY: build
build: clean
	@GOOS=linux go build -o ./oauth2 ./*.go

.PHONY: buildrelease
buildrelease:
	@GOOS=linux go build -ldflags="-s -w" -o ./oauth2 ./*.go
	@upx --brute ./oauth2

.PHONY: release
release: clean buildrelease
