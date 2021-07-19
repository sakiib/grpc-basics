package service

import (
	"errors"
	"github.com/jinzhu/copier"
	pb "github.com/sakiib/grpc-basics/gen.pb.go"
	"sync"
)

type BookStore interface {
	Save(book *pb.Book) error
}

type InMemoryBookStore struct {
	mutex sync.Mutex
	data  map[string]*pb.Book
}

func NewInMemoryBookStore() *InMemoryBookStore {
	return &InMemoryBookStore{
		data: make(map[string]*pb.Book),
	}
}

func (store *InMemoryBookStore) Save(book *pb.Book) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[book.Id] != nil {
		return errors.New("already exists")
	}

	other := &pb.Book{}
	err := copier.Copy(other, book)
	if err != nil {
		return err
	}
	store.data[other.Id] = other
	return nil
}
