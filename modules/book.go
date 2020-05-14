package modules

import "context"

type BookService interface {
	Create(ctx context.Context, book Book) (*Book, error)
	Books(ctx context.Context) ([]*Book, error)
	Book(ctx context.Context, id int) (*Book, error)
	Update(ctx context.Context, book *Book) (bool, error)
}


type Book struct {
	ID int
	Name string
	Author string
}
