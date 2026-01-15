#!/bin/bash

cd "$(dirname "$0")/.." || exit

# Очистка (опционально)
rm -rf ./internal/pb/models/*.pb.go
rm -rf ./internal/pb/users_api/*.pb.go ./internal/pb/users_api/*.pb.gw.go
rm -rf ./internal/pb/meals_api/*.pb.go ./internal/pb/meals_api/*.pb.gw.go
rm -rf ./internal/pb/swagger/users_api/*.json
rm -rf ./internal/pb/swagger/meals_api/*.json

mkdir -p ./internal/pb/models
mkdir -p ./internal/pb/users_api ./internal/pb/meals_api
mkdir -p ./internal/pb/swagger/users_api ./internal/pb/swagger/meals_api

# === Общие пути к зависимостям ===
PROTO_ROOT="./api"
GOOGLEAPIS_DIR=$(go list -f '{{ .Dir }}' -m github.com/grpc-ecosystem/grpc-gateway/v2)/protoc-gen-openapiv2
GOOGLE_GENPROTO_DIR=$(go list -f '{{ .Dir }}' -m google.golang.org/genproto)/googleapis

# Проверяем, что зависимости установлены
if [ ! -d "$GOOGLEAPIS_DIR" ]; then
  echo "protoc-gen-openapiv2 не найден. Выполните:"
  echo "go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest"
  exit 1
fi

if [ ! -d "$GOOGLE_GENPROTO_DIR" ]; then
  echo "googleapis не найден. Выполните:"
  echo "go mod tidy"
  exit 1
fi

protoc \
  -I "$PROTO_ROOT" \
  -I "$GOOGLE_GENPROTO_DIR" \
  --go_out=./internal/pb --go_opt=paths=source_relative \
  "$PROTO_ROOT/models/*.proto"

protoc \
  -I "$PROTO_ROOT" \
  -I "$PROTO_ROOT/google/api" \
  -I "$GOOGLE_GENPROTO_DIR" \
  --go_out=./internal/pb --go_opt=paths=source_relative \
  --go-grpc_out=./internal/pb --go-grpc_opt=paths=source_relative \
  --grpc-gateway_out=./internal/pb --grpc-gateway_opt=paths=source_relative,logtostderr=true \
  --openapiv2_out=./internal/pb/swagger --openapiv2_opt=logtostderr=true \
  "$PROTO_ROOT/users_api/users.proto"

protoc \
  -I "$PROTO_ROOT" \
  -I "$PROTO_ROOT/google/api" \
  -I "$GOOGLE_GENPROTO_DIR" \
  --go_out=./internal/pb --go_opt=paths=source_relative \
  --go-grpc_out=./internal/pb --go-grpc_opt=paths=source_relative \
  --grpc-gateway_out=./internal/pb --grpc-gateway_opt=paths=source_relative,logtostderr=true \
  --openapiv2_out=./internal/pb/swagger --openapiv2_opt=logtostderr=true \
  "$PROTO_ROOT/meal_api/meals.proto"

echo "Генерация завершена"