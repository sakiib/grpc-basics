package service

import (
	"context"
	pb "github.com/sakiib/grpc-basics/gen.pb.go"
)

type BookService struct {
	pb.UnimplementedBookServiceServer
	Store BookStore
}

func NewBookService(store BookStore) *BookService {
	return &BookService{pb.UnimplementedBookServiceServer{}, store}
}

func (b *BookService) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {
	book := req.GetBook()

	err := b.Store.Save(book)
	if err != nil {
		return nil, err
	}

	return &pb.CreateBookResponse{
		Id: book.Id,
	}, err
}

func (b *BookService) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.GetBookResponse, error) {
	id := req.GetId()
	book, err := b.Store.Get(id)
	if err != nil {
		return nil, err
	}
	return &pb.GetBookResponse{
		Book: book,
	}, nil
}
