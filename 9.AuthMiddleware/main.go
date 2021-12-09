package main

import (
	"fmt"
	"github.com/ewenliu/gin-example/9.AuthMiddleware/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func test1(c *gin.Context){
	fmt.Println("this is test1")
}

func login(c *gin.Context){
	auth.SetKeys("cookie", 1)
	c.String(http.StatusOK, "login successfully")
}

func dontNeedLogin(c *gin.Context){
	c.String(http.StatusOK, "dontNeedLogin")
}

func main() {
	r := gin.Default()
	testGroup := r.Group("/test")
	testGroup.Use(auth.LoginAuth())
	testGroup.GET("/info", func (c *gin.Context){
		c.String(http.StatusOK, "/test/info")
	})
	testGroup.POST("/test1", test1)
	r.POST("/login", login)
	r.GET("/no-login", dontNeedLogin)
	r.Run(":8888")
}
