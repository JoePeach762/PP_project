package bootstrap

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"

	userconsumer "github.com/JoePeach762/PP_project/user_service/internal/consumer/user"
	"github.com/JoePeach762/PP_project/user_service/internal/services/user"

	"github.com/JoePeach762/PP_project/user_service/internal/pb/users_api"

	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
	httpServer *http.Server
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) AppRun(
	ctx context.Context,
	userGRPC *user.GRPCServer,
	userConsumer *userconsumer.Consumer,
	httpPort int,
	grpcPort int,
) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go userConsumer.Consume(ctx)

	grpcAddr := fmt.Sprintf(":%d", grpcPort)
	errCh := make(chan error, 2)
	go func() {
		errCh <- s.runGRPCServer(grpcAddr, userGRPC)
	}()

	httpAddr := fmt.Sprintf(":%d", httpPort)
	go func() {
		errCh <- s.runGatewayServer(ctx, httpAddr, userGRPC)
	}()

	var runErr error
	select {
	case <-ctx.Done():
		slog.Info("Завершение работы...")
	case err := <-errCh:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			runErr = err
			slog.Error("Ошибка выполнения сервера", "error", err)
		}
		cancel()
	}

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if s.httpServer != nil {
		if err := s.httpServer.Shutdown(shutdownCtx); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("ошибка остановки gateway-сервера: %w", err)
		}
	}
	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
	}

	if runErr != nil {
		return fmt.Errorf("ошибка запуска сервера: %w", runErr)
	}

	return nil
}

func (s *Server) runGRPCServer(addr string, userGRPC *user.GRPCServer) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("не удалось начать прослушивание: %w", err)
	}

	s.grpcServer = grpc.NewServer()
	users_api.RegisterUserServiceServer(s.grpcServer, userGRPC)

	slog.Info("gRPC-сервер запущен", "addr", addr)
	return s.grpcServer.Serve(lis)
}

func (s *Server) runGatewayServer(ctx context.Context, httpAddr string, userGRPC *user.GRPCServer) error {
	r := chi.NewRouter()

	swaggerPath := os.Getenv("SWAGGER_PATH")
	if swaggerPath == "" {
		swaggerPath = "./internal/pb/swagger/users_api/users.swagger.json"
	}

	r.Get("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, swaggerPath)
	})
	r.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger.json"),
	))

	mux := runtime.NewServeMux()
	if err := users_api.RegisterUserServiceHandlerServer(ctx, mux, userGRPC); err != nil {
		return fmt.Errorf("не удалось зарегистрировать сервис пользователей: %w", err)
	}

	r.Mount("/", mux)

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	s.httpServer = &http.Server{
		Addr:    httpAddr,
		Handler: r,
	}

	slog.Info("HTTP/gRPC-Gateway сервер запущен", "addr", httpAddr)
	return s.httpServer.ListenAndServe()
}
