package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"github.com/guregu/dynamo"
	"github.com/pistatium/hikidashi_go/entities"
	"github.com/pistatium/hikidashi_go/infrastructures/dynamodb"
	"github.com/pistatium/hikidashi_go/repositories"
	"net/http"
)

type Backend int
const (
	DynamoDB Backend = iota
)

type ServerOptions struct {
	doInitialize bool
	initializationOptions interface{}
	backend Backend
	dynamodbEndpoint string
	dynamodbItemTableName string
}

type Server struct {
	itemRepo repositories.ItemRepository
}

func NewServer(ctx context.Context, opt ServerOptions) (*Server, error) {
	var repo repositories.ItemRepository
	switch opt.backend {
	case DynamoDB:
		sess, err := session.NewSession(&aws.Config{
			Endpoint: aws.String(opt.dynamodbEndpoint),
		})
		if err != nil {
			return nil, err
		}
		db := dynamo.New(sess, &aws.Config{})
		repo = dynamodb.NewItemRepository(db, opt.dynamodbItemTableName)
	default:
		return nil, fmt.Errorf("invalid backend: %v", opt.backend)
	}
	if opt.doInitialize {
		err := repo.Initialize(ctx, opt.initializationOptions)
		if err != nil {
			return nil, err
		}
	}
	return &Server{
		itemRepo: repo,
	}, nil
}

func (s *Server) healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

type CreateItemRequest struct {
	value string
	contentType string
}
func (s *Server) createItem(ctx *gin.Context) {
	var data CreateItemRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	item := entities.NewItem(
		ctx.Param("path"),
		data.value,
	)
	if data.contentType != "" {
		item.ContentType = data.contentType
	}
	err := s.itemRepo.Put(ctx, &item)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"item": item,
	})
}

func (s *Server) listItems(ctx *gin.Context) {
	items, err := s.itemRepo.List(ctx, "", "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}
