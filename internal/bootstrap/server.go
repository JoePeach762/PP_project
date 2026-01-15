package bootstrap

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"

	studentsapi "github.com/JoePeach762/PP_project/internal/api/student_service_api"
	userapi "github.com/JoePeach762/PP_project/internal/api/user_api"
	studentsconsumer "github.com/JoePeach762/PP_project/internal/consumer/students_Info_upsert_consumer"
	userconsumer "github.com/JoePeach762/PP_project/internal/consumer/user"

	"github.com/JoePeach762/PP_project/internal/pb/students_api"
	"github.com/JoePeach762/PP_project/internal/pb/users_api"
	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func AppRun(
	studentAPI *studentsapi.StudentServiceAPI,
	userAPI *userapi.UserHandler,
	studentConsumer *studentsconsumer.StudentInfoUpsertConsumer,
	userConsumer *userconsumer.Consumer,
) {
	// Запуск Kafka consumers
	go studentConsumer.Consume(context.Background())
	go userConsumer.Consume(context.Background())

	// Запуск gRPC сервера
	go func() {
		if err := runGRPCServer(studentAPI, userAPI); err != nil {
			panic(fmt.Errorf("failed to run gRPC server: %v", err))
		}
	}()

	// Запуск HTTP/gRPC-Gateway
	if err := runGatewayServer(); err != nil {
		panic(fmt.Errorf("failed to run gateway server: %v", err))
	}
}

func runGRPCServer(studentAPI *studentsapi.StudentServiceAPI, userAPI *userapi.UserHandler) error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	students_api.RegisterStudentsServiceServer(s, studentAPI)
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
	students_api.RegisterStudentsServiceHandlerFromEndpoint(ctx, mux, ":50051", opts)
	users_api.RegisterUserServiceHandlerFromEndpoint(ctx, mux, ":50051", opts)

	r.Mount("/", mux)

	slog.Info("gRPC-Gateway server listening on :8080")
	return http.ListenAndServe(":8080", r)
}
