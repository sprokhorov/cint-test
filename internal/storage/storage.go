package storage

import (
	"context"

	"github.com/sprokhorov/cint-test/pkg/schema"
)

// Storage is requiered to abstract the repository level.
type Storage interface {
	List(ctx context.Context) ([]*schema.Reminder, error)
	Create(ctx context.Context, cfg *schema.Reminder) error
	Ping() error
	Close() error
}
