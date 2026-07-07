.PHONY: run-api run-worker build

build:
	go build -o api .
	go build -o worker-bin ./worker

run-api:
	go run .

run-worker:
	go run ./worker
