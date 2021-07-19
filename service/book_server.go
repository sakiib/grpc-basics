package service

import (
	"context"
	"github.com/google/uuid"
	pb "github.com/sakiib/grpc-basics/gen.pb.go"
	"log"
)

type BookServer struct {
	Store BookStore
}

func NewBookServer(store *InMemoryBookStore) *BookServer {
	return &BookServer{}
}

func (server *BookServer) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {
	book := req.GetBook()
	log.Printf("received a create book request with id: %v, %v, %v, %v", book.Id, book.Name, book.Price, book.Author)

	if len(book.Id) > 0 {
		_, err := uuid.Parse(book.Id)
		if err != nil {
			return nil, err
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, err
		}
		book.Id = id.String()
	}

	err := server.Store.Save(book)
	if err != nil {
		return nil, err
	}

	log.Println("successfully saved book with: ", book.Id)
	res := &pb.CreateBookResponse{
		Id: book.Id,
	}
	return res, nil
}
