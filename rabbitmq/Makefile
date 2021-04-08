.PHONY: run docker.rabbitmq docker.worker docker.stop

run:
	go run main.go

docker.worker: docker.worker.build
	docker run --rm \
		--name dev-rabbitmq-worker \
		--network dev-network \
		dev-rabbitmq-worker

docker.rabbitmq:
	docker run --rm -d \
		--name dev-rabbitmq \
		--hostname dev-rabbitmq \
		--network dev-network \
		-v ${HOME}/dev-rabbitmq:/var/lib/rabbitmq \
		-v ${PWD}/configs/definitions.json:/opt/definitions.json:ro \
		-v ${PWD}/configs/rabbitmq.config:/etc/rabbitmq/rabbitmq.config:ro \
		-p 5672:5672 \
		-p 15672:15672 \
		rabbitmq:3-management

docker.worker.build:
	cd worker && docker build -t dev-rabbitmq-worker .

docker.stop:
	docker stop dev-rabbit docker.worker
	
docker.network:
	docker network inspect dev-network >/dev/null 2>&1 || \
	docker network create -d bridge dev-network
