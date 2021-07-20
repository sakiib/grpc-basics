package service

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	pb "github.com/sakiib/grpc-basics/gen.pb.go"
	"log"
	"sync"
)

type BookStore interface {
	Save(book *pb.Book) error
	Get(id string) (*pb.Book, error)
}

type InMemoryBookStore struct {
	mutex sync.RWMutex
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
		return errors.New("book already exists")
	}

	other := &pb.Book{}
	err := copier.Copy(other, book)
	if err != nil {
		return errors.New("failed to copy")
	}
	log.Println("successfully saved the book")
	store.data[book.Id] = other
	return nil
}

func (store *InMemoryBookStore) Get(id string) (*pb.Book, error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[id] == nil {
		return nil, fmt.Errorf("failed to get the book with id: %s", id)
	}
	return store.data[id], nil
}
