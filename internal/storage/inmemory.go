package storage

import (
	"context"

	"github.com/sprokhorov/cint-test/pkg/schema"
)

type InMemoryStorage struct {
	storage []*schema.Reminder
}

func NewInMemoryStorage() Storage {
	return &InMemoryStorage{storage: []*schema.Reminder{}}
}

func (ims *InMemoryStorage) List(ctx context.Context) ([]*schema.Reminder, error) {
	return ims.storage, nil
}

func (ims *InMemoryStorage) Create(ctx context.Context, rmd *schema.Reminder) error {
	ims.storage = append(ims.storage, rmd)
	return nil
}

func (ims *InMemoryStorage) Ping() error {
	return nil
}
func (ims *InMemoryStorage) Close() error {
	return nil
}
