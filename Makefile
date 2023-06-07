BIN := "./bin/app"
COMPOSE_PATH = deployments/docker-compose.yml


build:
	go build -v -o $(BIN) cmd/*.go

run:
	$(BIN)

up:
	docker compose -f $(COMPOSE_PATH) up -d --build --force-recreate

.PHONY: run build up