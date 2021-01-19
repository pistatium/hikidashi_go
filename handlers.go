package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"github.com/guregu/dynamo"
	"github.com/pistatium/hikidashi_go/infrastructures/dynamodb"
	"github.com/pistatium/hikidashi_go/repositories"
	"net/http"
)

type Backend int
const (
	DynamoDB Backend = iota
)

type ServerOptions struct {
	backend Backend
	dynamodbEndpoint string
}

type Server struct {
	itemRepo repositories.ItemRepository
}

func NewServer(opt ServerOptions) (*Server, error) {
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
		repo = dynamodb.NewItemRepository(db)
	default:
		return nil, fmt.Errorf("invalid backend: %v", opt.backend)
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
