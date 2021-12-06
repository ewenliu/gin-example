package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func helloWorld(c *gin.Context) {
	res := gin.H{
		"message": "hello world!",
	}
	c.JSON(http.StatusOK, res)
}

func main() {
	r := gin.Default()
	r.GET("/", helloWorld)
	_ = r.Run("127.0.0.1:8888")
}
