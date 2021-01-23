package dynamodb

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/guregu/dynamo"
	"github.com/pistatium/hikidashi_go/entities"
	"github.com/pistatium/hikidashi_go/repositories"
	"strings"
	"time"
)

type ItemRepository struct {
	db *dynamo.DB
	tableName string
}

func (i ItemRepository) Initialize(ctx context.Context, options interface{}) error {
	_, err := i.db.Table(i.tableName).Describe().RunWithContext(ctx)
	if err == nil {
		return nil
	}
	if awsErr, ok := err.(awserr.Error); ok {
		if awsErr.Code() == dynamodb.ErrCodeResourceNotFoundException {
			input := i.db.CreateTable(i.tableName, itemTable{}).
				Provision(1, 1)
			err = input.RunWithContext(ctx)
			return err
		}
		return awsErr
	}
	return nil
}

func NewItemRepository(db *dynamo.DB, tableName string) repositories.ItemRepository {
	return &ItemRepository{
		db: db,
		tableName: tableName,
	}
}

func (i ItemRepository) GetFromPath(ctx context.Context, path string) (item *entities.Item, err error) {
	table := i.db.Table(i.tableName)
	err = table.Get("Path", path).OneWithContext(ctx, item)
	return
}

func (i ItemRepository) Put(ctx context.Context, item *entities.Item) (err error) {
	group, groupPath := splitPath(item.Path)
	it := itemTable{
		Path:        item.Path,
		Group:       group,
		GroupPath:   groupPath,
		Value:       item.Value,
		ContentType: item.ContentType,
		UpdatedAt:   time.Now(),
	}
	table := i.db.Table(i.tableName)
	err = table.Put(it).RunWithContext(ctx)
	return
}

func (i ItemRepository) List(ctx context.Context, directory string, nextCursor string) (items *[]entities.Item, err error) {
	table := i.db.Table(i.tableName)
	tableItems := make([]itemTable, 0)
	results := make([]entities.Item, 0)
	if directory == "" {
		err = table.Get("Group", directory).AllWithContext(ctx, tableItems)
		if err != nil {
			return
		}
		for _, ti := range tableItems {
			results = append(results, entities.Item{
				Path:        ti.Path,
				Value:       ti.Value,
				ContentType: ti.ContentType,
			})
		}
	}
	err = table.Scan().AllWithContext(ctx, items)
	items = &results
	return
}

func splitPath(path string) (group string, groupPath string) {
	ss := strings.Split(path, "/")
	group = ss[0]
	groupPath = strings.Join(ss[1:], "/")
	return
}
