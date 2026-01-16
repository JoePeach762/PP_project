#!/bin/bash

cd "$(dirname "$0")/.." || exit

rm -rf ./user_service/internal/pb/models/*.pb.go
rm -rf ./meal_service/internal/pb/models/*.pb.go
rm -rf ./user_service/internal/pb/users_api/*.pb.go ./user_service/internal/pb/users_api/*.pb.gw.go
rm -rf ./meal_service/internal/pb/meals_api/*.pb.go ./meal_service/internal/pb/meals_api/*.pb.gw.go
rm -rf ./user_service/internal/pb/swagger/users_api/*.json
rm -rf ./meal_service/internal/pb/swagger/meals_api/*.json

mkdir -p ./user_service/internal/pb/models
mkdir -p ./meal_service/internal/pb/models
mkdir -p ./user_service/internal/pb/users_api ./meal_service/internal/pb/meals_api
mkdir -p ./user_service/internal/pb/swagger/users_api ./meal_service/internal/pb/swagger/meals_api

PROTO_ROOT="./api"

protoc \
  -I "$PROTO_ROOT" \
  --go_out=./user_service/internal/pb --go_opt=paths=source_relative \
  "$PROTO_ROOT/models/user_models/*.proto"

protoc \
  -I "$PROTO_ROOT" \
  --go_out=./meal_service/internal/pb --go_opt=paths=source_relative \
  "$PROTO_ROOT/models/meal_models/*.proto"

protoc \
  -I "$PROTO_ROOT" \
  -I "$PROTO_ROOT/google/api" \
  --go_out=./user_service/internal/pb --go_opt=paths=source_relative \
  --go-grpc_out=./user_service/internal/pb --go-grpc_opt=paths=source_relative \
  --grpc-gateway_out=./user_service/internal/pb --grpc-gateway_opt=paths=source_relative,logtostderr=true \
  --openapiv2_out=./user_service/internal/pb/swagger --openapiv2_opt=logtostderr=true \
  "$PROTO_ROOT/users_api/users.proto"

protoc \
  -I "$PROTO_ROOT" \
  -I "$PROTO_ROOT/google/api" \
  --go_out=./meal_service/internal/pb --go_opt=paths=source_relative \
  --go-grpc_out=./meal_service/internal/pb --go-grpc_opt=paths=source_relative \
  --grpc-gateway_out=./meal_service/internal/pb --grpc-gateway_opt=paths=source_relative,logtostderr=true \
  --openapiv2_out=./meal_service/internal/pb/swagger --openapiv2_opt=logtostderr=true \
  "$PROTO_ROOT/meals_api/meals.proto"

echo "Генерация завершена"