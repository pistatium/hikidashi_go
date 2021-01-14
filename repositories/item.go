package repositories

import (
	"context"
	"github.com/pistatium/hikidashi_go/entities"
)

type ItemRepository interface {
	GetFromPath(ctx context.Context, path string) (entities.Item, error)
	Put(ctx context.Context, item entities.Item) error
}
