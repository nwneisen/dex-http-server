package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc/grpclog"
)

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcServerEndpoint := getEnv("GRPC_SERVER_ENDPOINT", "localhost:9090")
	fmt.Printf("\033[32m%s\033[0m\n", "gRPC server endpoint registered: "+grpcServerEndpoint)

	httpServerPort := getEnv("HTTP_SERVER_PORT", "8081")
	fmt.Printf("\033[32m%s\033[0m\n", "HTTP server listening on port: "+httpServerPort)

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	// opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	// err := gw.RegisterYourServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	// if err != nil {
	// 	return err
	// }

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":"+httpServerPort, mux)
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		grpclog.Fatal(err)
	}
}
