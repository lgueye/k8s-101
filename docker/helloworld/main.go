package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	name := "World"
	data, err := ioutil.ReadFile("config/name")
	if err == nil {
		name = string(data)
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello " + name + "!")
	})
	r.GET("/sickz", func (c *gin.Context) {
		c.JSON(http.StatusInternalServerError, "I'm sick!")
	})
	r.GET("/healthz", func (c *gin.Context) {
		c.JSON(http.StatusOK, "I'm healthy!")
	})
	_ = r.Run(":8080")
}
