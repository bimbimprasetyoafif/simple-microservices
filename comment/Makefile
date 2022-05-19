export GO111MODULE=on

GO_CMD=go
GO_TEST_VERBOSE=$(GO_CMD) test ./... -v
GO_TEST_INTEGRATION==$(GO_CMD) test -tags=integration
GO_GET=$(GO_CMD) get -v ./...
GO_BUILD=$(GO_CMD) build

DOCKER_COMPOSE=docker-compose

SERVER_BIN=main

dependency:
	@echo "==> Running Install Dependency ...";
	@$(GO_GET)

tests:
	@echo "==> Running tests ...";
	@$(GO_TEST_VERBOSE)

build:
	@echo "==> Building server ...";
	@$(GO_BUILD) -o $(SERVER_BIN) -ldflags "-w -s" ./main.go || exit 1;
	@chmod 755 $(SERVER_BIN)

run:
	./$(SERVER_BIN)

e2e:
	@echo "==> Start Integration Test ...";
	@echo "==> Init db ...";
	@$(DOCKER_COMPOSE) up -d  --remove-orphans
	@sleep 3
	@echo "==> run test ...";
	@$(GO_TEST_VERBOSE) ./testing -v -count=1
	@echo "==> remove db ...";
	@$(DOCKER_COMPOSE) down
	@echo "==> integration testing done ...";
