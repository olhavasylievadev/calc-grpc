//Implementation of calculator server.
//By default runs on port ":8081"
//Provides access to GRPC server methods.

package main

import (
	"github.com/olhavasylievadev/calc-grpc/pkg/api"
	"github.com/olhavasylievadev/calc-grpc/pkg/calculator"
	"google.golang.org/grpc"
	"log"
	"net"
)

// The server is working via GRPC calls
func main() {
	s := grpc.NewServer()
	srv := &calculator.GRPCServer{}
	api.RegisterCalculatingServiceServer(s, srv)

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
