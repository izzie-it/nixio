package nixio

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/izzie-it/nixio/log"
	"go.elastic.co/apm/module/apmgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func NewServer() *grpc.Server {
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(
		apmgrpc.NewUnaryServerInterceptor(apmgrpc.WithRecovery()),
	))

	return grpcServer
}

func Run(grpcServer *grpc.Server, serviceName string, port int) error {
	healthServer := health.NewServer()
	healthpb.RegisterHealthServer(grpcServer, healthServer)
	healthServer.SetServingStatus(serviceName, healthpb.HealthCheckResponse_NOT_SERVING)
	go func() {
		for {
			status := healthpb.HealthCheckResponse_SERVING

			healthServer.SetServingStatus(serviceName, status)

			time.Sleep(500 * time.Millisecond)
		}
	}()

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		// log.Fatalf("Failed to listing: %s\n", err)
		return err
	}

	go func() {
		log.Infof("Starting server on port %d\n", port)
		if err := grpcServer.Serve(lis); err != nil {
			// log.Fatalf("Failed to serve: %s\n", err.Error())
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutting down server...")

	grpcServer.GracefulStop()

	log.Info("Server exiting")
	return nil
}
