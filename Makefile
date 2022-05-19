DOCKER_COMPOSE=docker-compose

run:
	@echo "==> Server running ...";
	@$(DOCKER_COMPOSE) up -d  --remove-orphans
	@echo "==> Running REST at :8000 ...";
	@echo "==> Please using make stop to stop server ...";

stop:
	@$(DOCKER_COMPOSE) down
	@echo "==> Server down ...";
