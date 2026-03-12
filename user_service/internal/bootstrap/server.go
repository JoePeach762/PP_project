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
	"google.golang.org/grpc/credentials/insecure"
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
	grpcReady := make(chan struct{})
	errCh := make(chan error, 2)
	go func() {
		errCh <- s.runGRPCServer(grpcAddr, userGRPC, grpcReady)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errCh:
		if err != nil {
			return fmt.Errorf("gRPC server failed: %w", err)
		}
		return nil
	case <-grpcReady:
	}

	httpAddr := fmt.Sprintf(":%d", httpPort)
	go func() {
		errCh <- s.runGatewayServer(ctx, httpAddr, grpcAddr)
	}()

	var runErr error
	select {
	case <-ctx.Done():
		slog.Info("Shutting down...")
	case err := <-errCh:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			runErr = err
			slog.Error("Server runtime failed", "error", err)
		}
		cancel()
	}

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if s.httpServer != nil {
		if err := s.httpServer.Shutdown(shutdownCtx); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("shutdown gateway server: %w", err)
		}
	}
	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
	}

	if runErr != nil {
		return fmt.Errorf("server run failed: %w", runErr)
	}

	return nil
}

func (s *Server) runGRPCServer(addr string, userGRPC *user.GRPCServer, ready chan<- struct{}) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	s.grpcServer = grpc.NewServer()
	users_api.RegisterUserServiceServer(s.grpcServer, userGRPC)
	close(ready)

	slog.Info("gRPC server listening", "addr", addr)
	return s.grpcServer.Serve(lis)
}

func (s *Server) runGatewayServer(ctx context.Context, httpAddr, grpcAddr string) error {
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
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	if err := users_api.RegisterUserServiceHandlerFromEndpoint(ctx, mux, grpcAddr, opts); err != nil {
		return fmt.Errorf("failed to register user service: %w", err)
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

	slog.Info("HTTP/gRPC-Gateway server listening", "addr", httpAddr)
	return s.httpServer.ListenAndServe()
}
