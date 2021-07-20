package main

import (
	"flag"
	"fmt"
	pb "github.com/sakiib/grpc-basics/gen.pb.go"
	"github.com/sakiib/grpc-basics/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	port := flag.Int("port", 0, "server port")
	flag.Parse()

	log.Printf("start server on port: %d", *port)

	grpcServer := grpc.NewServer()
	bookServer := service.NewBookService()
	pb.RegisterBookServiceServer(grpcServer, bookServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
