package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	opt := ServerOptions{
		backend: DynamoDB,
	}
	s, err := NewServer(opt)
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	router.GET("/health", s.healthCheck)
	router.GET("/items", s.listItems)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	err = router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		fmt.Println(err)
	}
}


