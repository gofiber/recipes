.PHONY: start
start:
	docker-compose up -d --build

.PHONY: stop
stop:
	docker-compose down

.PHONY: test
test:
	sh ./scripts/e2e-testing.sh
