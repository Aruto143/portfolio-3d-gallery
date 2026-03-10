COMPOSE_FILE=infra/docker/docker-compose.yml

up:
	docker compose -f $(COMPOSE_FILE) up -d --build

down:
	docker compose -f $(COMPOSE_FILE) down

logs:
	docker compose -f $(COMPOSE_FILE) logs -f

ps:
	docker compose -f $(COMPOSE_FILE) ps

restart:
	docker compose -f $(COMPOSE_FILE) down
	docker compose -f $(COMPOSE_FILE) up -d --build

backend-logs:
	docker compose -f $(COMPOSE_FILE) logs -f backend

frontend-logs:
	docker compose -f $(COMPOSE_FILE) logs -f frontend

db-logs:
	docker compose -f $(COMPOSE_FILE) logs -f db