package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	var env Env
	err := envconfig.Process("", &env)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%+v", env)
	opt := ServerOptions{
		backend: DynamoDB,
		dynamodbEndpoint: env.DynamodbEndpointUrl,
		dynamodbItemTableName: env.DynamodbItemTableName,
		doInitialize: env.RequireInitializeRepository,
	}
	s, err := NewServer(context.Background(), opt)
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	{
		router.GET("/health", s.healthCheck)
	}
	itemRouter := router.Group("/items")
	{
		itemRouter.POST("/:path", s.createItem)

		itemRouter.GET("/:path", s.listItems)
		itemRouter.GET("", s.listItems)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	err = router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		fmt.Println(err)
	}
}


