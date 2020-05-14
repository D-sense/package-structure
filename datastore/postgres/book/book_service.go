package book

import (
	"context"
	"time"
)


type BookService struct{}

func Create(ctx context.Context, book Book) (modules.book, error) {
	var track []*modules.Book
	err := ts.client.db.Select("id, counter, original_resource, metadata").Where("bpm is null AND deleted_at is null").Order("counter desc").Find(&track).Error
	return track, err
}

func (ts BookService) Books(ctx context.Context, excludeIds []string) ([]*modules.Books, error) {
	var track []*modules.Book
	err := ts.client.db.Select("id, counter, original_resource, metadata").Where("bpm is null AND deleted_at is null").Order("counter desc").Find(&track).Error
	return track, err
}

func (ts BookService) Book(ctx context.Context, id int) (*modules.Book, error) {
	var track []*modules.Book
	err := ts.client.db.Select("id, counter, original_resource, metadata").Where("bpm is null AND deleted_at is null").Order("counter desc").Find(&track).Error
	return track, err
}

func (ts BookService) Update(ctx context.Context, book *modules.Book) error {
	track.UpdatedAt = time.Now()
	return ts.client.db.Save(&track).Error
}