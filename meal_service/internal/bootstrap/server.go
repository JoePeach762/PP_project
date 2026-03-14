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

	mealconsumer "github.com/JoePeach762/PP_project/meal_service/internal/consumer/meal"
	"github.com/JoePeach762/PP_project/meal_service/internal/services/meal"

	"github.com/JoePeach762/PP_project/meal_service/internal/pb/meals_api"

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
	mealGRPC *meal.GRPCServer,
	mealConsumer *mealconsumer.Consumer,
	httpPort int,
	grpcPort int,
) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go mealConsumer.Consume(ctx)

	grpcAddr := fmt.Sprintf(":%d", grpcPort)
	errCh := make(chan error, 2)
	go func() {
		errCh <- s.runGRPCServer(grpcAddr, mealGRPC)
	}()

	httpAddr := fmt.Sprintf(":%d", httpPort)
	go func() {
		errCh <- s.runGatewayServer(ctx, httpAddr, mealGRPC)
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

func (s *Server) runGRPCServer(addr string, mealGRPC *meal.GRPCServer) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("не удалось начать прослушивание: %w", err)
	}

	s.grpcServer = grpc.NewServer()
	meals_api.RegisterMealServiceServer(s.grpcServer, mealGRPC)

	slog.Info("gRPC-сервер запущен", "addr", addr)
	return s.grpcServer.Serve(lis)
}

func (s *Server) runGatewayServer(ctx context.Context, httpAddr string, mealGRPC *meal.GRPCServer) error {
	r := chi.NewRouter()

	swaggerPath := os.Getenv("SWAGGER_PATH")
	if swaggerPath == "" {
		swaggerPath = "./internal/pb/swagger/meals_api/meals.swagger.json"
	}

	r.Get("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, swaggerPath)
	})
	r.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger.json"),
	))

	mux := runtime.NewServeMux()
	if err := meals_api.RegisterMealServiceHandlerServer(ctx, mux, mealGRPC); err != nil {
		return fmt.Errorf("не удалось зарегистрировать сервис приемов пищи: %w", err)
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
