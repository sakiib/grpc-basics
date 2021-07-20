package service

import (
	"context"
	"fmt"
	pb "github.com/sakiib/grpc-basics/gen.pb.go"
	"github.com/sakiib/grpc-basics/helpers"
	"log"
)

type BookService struct {
	pb.UnimplementedBookServiceServer
}

func NewBookService() *BookService {
	return &BookService{}
}

func (b *BookService) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {
	book := req.Book
	author := req.Book.Author

	log.Println("book details: ", book.Id, book.Name, book.Price, book.Isbn, author.FirstName, author.LastName)

	res := &pb.CreateBookResponse{
		Id: fmt.Sprintf("%s-%s", book.Id, book.Name),
	}

	return res, nil
}

func (b *BookService) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.GetBookResponse, error) {
	res := &pb.GetBookResponse{
		Book: helpers.RandBook(),
	}
	return res, nil
}
