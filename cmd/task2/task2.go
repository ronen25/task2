package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/ronen25/task2/instrumentation"
	"github.com/ronen25/task2/service"
	"github.com/ronen25/task2/service/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func httpServerRoutine(serverPort string, doneChannel <-chan os.Signal) {
	service := service.NewHTTPService()

	// Initialize router
	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router))

	// Register instrumentation
	ins := instrumentation.NewInstrumentation()
	ins.RegisterInstrumentation()

	router.Handle("/metrics", ins.HTTPHandler)
	router.HandleFunc("/", service.QueryPrintingHandler)

	// Listen on provided HTTP port
	bindAddress := fmt.Sprintf(":%s", serverPort)
	server := &http.Server{Addr: bindAddress, Handler: router}

	go func() {
		log.Printf("HTTP listening on address %s", bindAddress)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Error setting up HTTP server: %v", err)
		}
	}()

	<-doneChannel
	log.Printf("Exiting HTTP server...")

	// Exit handler
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error shutting down HTTP server: %v", err)
	}
}

func grpcServerRoutine(serverPort string, doneChannel <-chan os.Signal) {
	bindAddr := fmt.Sprintf(":%s", serverPort)
	lis, err := net.Listen("tcp", bindAddr)
	if err != nil {
		log.Fatalf("Error setting up GRPC endpoint: %v", err)
	}

	grpcServer := grpc.NewServer()
	protos.RegisterQueryPrinterGRPCServer(grpcServer, service.NewQueryPrinterGRPCService())
	reflection.Register(grpcServer)
	go func() {
		log.Printf("GRPC listening on address %s", bindAddr)
		grpcServer.Serve(lis)
	}()

	<-doneChannel
	log.Printf("Exiting GRPC server...")

	// Exit handler
	grpcServer.GracefulStop()
}

func main() {
	signalChannel := make(chan os.Signal, 1)

	portVars := map[string]string{
		"HTTP_PORT": "",
		"GRPC_PORT": "",
	}

	// Get ports from the environment
	for key := range portVars {
		if value, exists := os.LookupEnv(key); exists {
			portVars[key] = value
		} else {
			// Variable unavailable = error
			log.Fatalf("Error: Cannot find required environment variable `%s`", key)
		}
	}

	// Register SIGINT signal handler. We use this to quit the app properly.
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	// Initialize both services - one on HTTP, the other on GRPC
	go httpServerRoutine(portVars["HTTP_PORT"], signalChannel)
	go grpcServerRoutine(portVars["GRPC_PORT"], signalChannel)

	// App quits when both servers are done
	<-signalChannel
}
