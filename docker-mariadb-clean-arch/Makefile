.PHONY: start
start:
	docker-compose up -d --build

.PHONY: stop
stop:
	docker-compose rm -v --force --stop
	docker image rm docker-mariadb-clean-arch:latest

.PHONY: test
test:
	sh ./scripts/e2e-testing.sh
