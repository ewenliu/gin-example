package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func helloWorld(c *gin.Context) {
	name := c.PostForm("name")
	age := c.PostForm("age")
	res := fmt.Sprintf("%s is %s years old", name, age)
	c.String(http.StatusOK, res)
}

func main() {
	r := gin.Default()
	r.POST("/", helloWorld)
	// 使用postman在body的form-data中传入对应的参数即可
	_ = r.Run("127.0.0.1:8888")
}