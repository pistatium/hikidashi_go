package dynamodb

import (
	"context"
	"github.com/guregu/dynamo"
	"github.com/pistatium/hikidashi_go/entities"
	"github.com/pistatium/hikidashi_go/repositories"
	"strings"
	"time"
)

const ItemTableName = "Items"

type ItemRepository struct {
	db *dynamo.DB
}

func NewItemRepository(db *dynamo.DB) repositories.ItemRepository {
	return &ItemRepository{
		db: db,
	}
}

func (i ItemRepository) GetFromPath(ctx context.Context, path string) (item *entities.Item, err error) {
	table := i.db.Table(ItemTableName)
	err = table.Get("Path", path).One(item)
	return
}

func (i ItemRepository) Put(ctx context.Context, item entities.Item) (err error) {
	group, groupPath := splitPath(item.Path)
	it := itemTable{
		Path:        item.Path,
		Group:       group,
		GroupPath:   groupPath,
		Value:       item.Value,
		ContentType: item.ContentType,
		UpdatedAt:   time.Now(),
	}
	table := i.db.Table(ItemTableName)
	err = table.Put(it).Run()
	return
}

func (i ItemRepository) List(ctx context.Context, directory string, nextCursor string) (items []entities.Item, err error) {
	panic("implement me")
}

func splitPath(path string) (group string, groupPath string) {
	ss := strings.Split(path, "/")
	group = ss[0]
	groupPath = strings.Join(ss[1:], "/")
	return
}
