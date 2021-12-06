package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func helloWorld(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	res := fmt.Sprintf("%s is %s years old", name, age)
	c.String(http.StatusOK, res)
}

func main() {
	r := gin.Default()
	r.GET("/", helloWorld)
	// 访问http://127.0.0.1:8888/?name=liuyiwen&&age=18
	_ = r.Run("127.0.0.1:8888")
}
