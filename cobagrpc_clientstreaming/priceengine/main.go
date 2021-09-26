package main

import (
	"flag"
	"fmt"
	"github.com/jonnyGit81/coba/cobagrpc_clientstreaming/price_api"
	"github.com/jonnyGit81/coba/cobagrpc_clientstreaming/priceengine/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 9111, "Port on which gRPC server should listen TCP conn.")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	price_api.RegisterAPIServer(grpcServer, &server.PriceServer{})
	grpcServer.Serve(lis)
	log.Printf("Initializing gRPC server on port %d", *port)
}
