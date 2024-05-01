package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func health(c *gin.Context) {
	c.JSON(http.StatusOK, "health")
}

func stress(c *gin.Context) {
	c.JSON(http.StatusOK, "stress")
}

func main() {
	port := os.Getenv("BIND-PORT")
	router := gin.Default()
	router.GET("/", health)
	router.GET("/stress", stress)
	router.Run(fmt.Sprintf(":%s", port))
}
