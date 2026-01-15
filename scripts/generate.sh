#!/bin/bash

cd "$(dirname "$0")/.." || exit

# Очистка
rm -rf ./internal/pb/models/*.pb.go
rm -rf ./internal/pb/users_api/*.pb.go ./internal/pb/users_api/*.pb.gw.go
rm -rf ./internal/pb/meals_api/*.pb.go ./internal/pb/meals_api/*.pb.gw.go
rm -rf ./internal/pb/swagger/users_api/*.json
rm -rf ./internal/pb/swagger/meals_api/*.json

mkdir -p ./internal/pb/models
mkdir -p ./internal/pb/users_api ./internal/pb/meals_api
mkdir -p ./internal/pb/swagger/users_api ./internal/pb/swagger/meals_api

PROTO_ROOT="./api"

# === Генерация моделей ===
protoc \
  -I "$PROTO_ROOT" \
  --go_out=./internal/pb --go_opt=paths=source_relative \
  "$PROTO_ROOT/models/*.proto"

# === Генерация Users API ===
protoc \
  -I "$PROTO_ROOT" \
  -I "$PROTO_ROOT/google/api" \
  --go_out=./internal/pb --go_opt=paths=source_relative \
  --go-grpc_out=./internal/pb --go-grpc_opt=paths=source_relative \
  --grpc-gateway_out=./internal/pb --grpc-gateway_opt=paths=source_relative,logtostderr=true \
  --openapiv2_out=./internal/pb/swagger --openapiv2_opt=logtostderr=true \
  "$PROTO_ROOT/users_api/users.proto"

# === Генерация Meals API ===
protoc \
  -I "$PROTO_ROOT" \
  -I "$PROTO_ROOT/google/api" \
  --go_out=./internal/pb --go_opt=paths=source_relative \
  --go-grpc_out=./internal/pb --go-grpc_opt=paths=source_relative \
  --grpc-gateway_out=./internal/pb --grpc-gateway_opt=paths=source_relative,logtostderr=true \
  --openapiv2_out=./internal/pb/swagger --openapiv2_opt=logtostderr=true \
  "$PROTO_ROOT/meals_api/meals.proto"

echo "Генерация завершена"