package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func main() {
	logfile, err := os.OpenFile("/logs/out", os.O_APPEND | os.O_WRONLY, 0644)
	if os.IsNotExist(err) {
		logfile, err = os.Create("/logs/out")
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}
	defer logfile.Close()
	logfile.WriteString("======================\n")

	name := "World"
	data, err := ioutil.ReadFile("config/name")
	if err == nil {
		name = string(data)
	}

	r := gin.Default()
	var count int64
	count = 0
	r.GET("/", func(c *gin.Context) {
		logfile.WriteString(strconv.FormatInt(count, 10) + "/ GET" + c.FullPath() + "\n")
		count++
		c.JSON(http.StatusOK, "Hello " + name + "!")
	})
	r.GET("/version", func (c *gin.Context) {
		logfile.WriteString(strconv.FormatInt(count, 10) + "/ GET" + c.FullPath() + "\n")
		count++
		c.JSON(http.StatusOK, "1.0")
	})
	r.GET("/sickz", func (c *gin.Context) {
		logfile.WriteString(strconv.FormatInt(count, 10) + "/ GET" + c.FullPath() + "\n")
		count++
		c.JSON(http.StatusInternalServerError, "I'm sick!")
	})
	r.GET("/healthz", func (c *gin.Context) {
		logfile.WriteString(strconv.FormatInt(count, 10) + "/ GET" + c.FullPath() + "\n")
		count++
		c.JSON(http.StatusOK, "I'm healthy!")
	})
	_ = r.Run(":8080")
}
