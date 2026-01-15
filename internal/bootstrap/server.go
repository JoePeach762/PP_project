package bootstrap

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"

	userapi "github.com/JoePeach762/PP_project/internal/api/user_api"
	userconsumer "github.com/JoePeach762/PP_project/internal/consumer/user"

	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func AppRun(
	userAPI *userapi.UserAPI,
	userConsumer *userconsumer.Consumer,
) {
	// Запуск Kafka consumers
	go userConsumer.Consume(context.Background())

	// Запуск gRPC сервера
	go func() {
		if err := runGRPCServer(userAPI); err != nil {
			panic(fmt.Errorf("failed to run gRPC server: %v", err))
		}
	}()

	// Запуск HTTP/gRPC-Gateway
	if err := runGatewayServer(); err != nil {
		panic(fmt.Errorf("failed to run gateway server: %v", err))
	}
}

func runGRPCServer(userAPI *userapi.UserAPI) error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	users_api.RegisterUserServiceServer(s, userAPI) // ← новая строка

	slog.Info("gRPC server listening on :50051")
	return s.Serve(lis)
}

func runGatewayServer() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	swaggerPath := os.Getenv("SWAGGER_PATH")
	if swaggerPath == "" {
		swaggerPath = "./internal/pb/swagger/students_api/students.swagger.json"
	}

	r := chi.NewRouter()
	r.Get("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, swaggerPath)
	})
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("/swagger.json")))

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Регистрация обоих сервисов
	users_api.RegisterUserServiceHandlerFromEndpoint(ctx, mux, ":50051", opts)

	r.Mount("/", mux)

	slog.Info("gRPC-Gateway server listening on :8080")
	return http.ListenAndServe(":8080", r)
}
