.PHONY: generate-api
generate-api:
	@./scripts/generate.sh

.PHONY: up
up:
	docker-compose up -d --build

.PHONY: down
down:
	docker-compose down

.PHONY: test-user-service
test-user-service:
	go -C user_service test ./internal/services/user/...

.PHONY: cov-user-service
cov-user-service:
	go -C user_service test -cover ./internal/services/user/...

.PHONY: mock-user-service
mock-user-service:
	cd user_service && mockery --config ../.mockery.yaml
