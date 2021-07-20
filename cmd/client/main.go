package main

import (
	"context"
	"flag"
	pb "github.com/sakiib/grpc-basics/gen.pb.go"
	"github.com/sakiib/grpc-basics/helpers"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server: %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	bookServer := pb.NewBookServiceClient(conn)
	book := helpers.RandBook()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := bookServer.CreateBook(ctx, &pb.CreateBookRequest{
		Book: book,
	})
	if err != nil {
		log.Println(err)
	}

	log.Println("response: ", res)
}
