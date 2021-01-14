package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func main() {
	router := gin.Default()
	router.GET("/health", healthCheck)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		fmt.Println(err)
	}
}
