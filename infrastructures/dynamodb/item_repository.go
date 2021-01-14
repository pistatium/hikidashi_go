package dynamodb

import (
	"context"
	"github.com/pistatium/hikidashi_go/entities"
)

type ItemRepository struct {}

func (i ItemRepository) GetFromPath(ctx context.Context, path string) (entities.Item, error) {
	panic("implement me")
}

func (i ItemRepository) Put(ctx context.Context, item entities.Item) error {
	panic("implement me")
}
