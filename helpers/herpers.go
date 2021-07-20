package helpers

import (
	pb "github.com/sakiib/grpc-basics/gen.pb.go"
)

func RandBook() *pb.Book {
	book := &pb.Book{
		Name:  "book-1",
		Id:    "id-12345",
		Isbn:  "isbn-12345",
		Price: 2.5,
		Author: &pb.Author{
			FirstName: "sakib",
			LastName:  "alamin",
		},
	}
	return book
}
