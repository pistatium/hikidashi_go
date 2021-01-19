package repositories

import (
	"context"
	"github.com/pistatium/hikidashi_go/entities"
)

type ItemRepository interface {
	Initialize(ctx context.Context, options interface{}) error
	GetFromPath(ctx context.Context, path string) (item *entities.Item, err error)
	Put(ctx context.Context, item *entities.Item) (err error)
	List(ctx context.Context, directory string, nextCursor string) (items *[]entities.Item, err error)
}
