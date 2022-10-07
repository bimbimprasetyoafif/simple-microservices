DOCKER_COMPOSE=docker-compose

run:
	@echo "==> Server running ...";
	@$(DOCKER_COMPOSE) up -d  --remove-orphans
	@echo "==> Running REST at :8000 ...";
	@echo "==> Please using make stop to stop server ...";

stop:
	@$(DOCKER_COMPOSE) down
	@echo "==> Server down ...";

build-comment:
	export GO111MODULE=on;
	GO_ENABLED=0 go build -a -installsuffix cgo=1 -o ./comment.app ./comment/main.go

build-org:
	export GO111MODULE=on;
	GO_ENABLED=0 go build -a -installsuffix cgo=1 -o ./org.app ./org/main.go

test:
	export GO111MODULE on; \
	go test ./... -cover -vet=all -v -short -covermode=count -coverprofile=cover.out > test.txt
