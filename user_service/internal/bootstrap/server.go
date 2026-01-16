package bootstrap

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	userGRPC *user.GRPCServer,
	userConsumer *userconsumer.Consumer,
) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go userConsumer.Consume(ctx)

	grpcAddr := ":50051"
	go func() {
		if err := s.runGRPCServer(grpcAddr, userGRPC); err != nil {
			slog.Error("gRPC server failed", "error", err)
			cancel()
		}
	}()

	time.Sleep(100 * time.Millisecond)

	httpAddr := ":8080"
	if err := s.runGatewayServer(ctx, httpAddr, grpcAddr); err != nil {
		return fmt.Errorf("gateway server failed: %w", err)
	}

	<-sigChan
	slog.Info("Shutting down...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if s.httpServer != nil {
		s.httpServer.Shutdown(shutdownCtx)
	}
	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
	}

	return nil
}

func (s *Server) runGRPCServer(addr string, userGRPC *user.GRPCServer) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	s.grpcServer = grpc.NewServer()
	users_api.RegisterUserServiceServer(s.grpcServer, userGRPC)

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
