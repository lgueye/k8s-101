package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getHello(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World!")
}

func main() {
	r := gin.Default()
	r.GET("/", getHello)
	r.Run(":8080")
}
