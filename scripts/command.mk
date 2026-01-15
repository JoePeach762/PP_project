.PHONY: generate-api
generate-api:
	@./scripts/generate.sh

.PHONY: up
up:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down

.PHONY: cov
cov:
	go test -cover ./... 

.PHONY: mock
mock:
	mockery
