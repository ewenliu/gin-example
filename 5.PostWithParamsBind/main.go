package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

type person struct {
	Name string `form:"name" binding:"required"`
	Age  int    `form:"age,default=20" binding:"min=0,max=99"`
	BankBalance float64 `form:"bankBalance,default=999999.99"`
	HasChild bool `form:"hasChild, default=false"`
	Birthday time.Time `form:"birthday,default=1994-11-20" time_format:"2006-01-02"`
}

func helloWorld(c *gin.Context) {
	p := &person{}
	if err := c.ShouldBind(p); err != nil {
		log.Println(err)
		c.String(http.StatusOK, "error params")
	} else {
		res := fmt.Sprintf("%s is %s years old, borned in %s", p.Name, strconv.Itoa(p.Age), p.Birthday)
		c.String(http.StatusOK, res)
	}
}

func main() {
	r := gin.Default()
	r.POST("/", helloWorld)
	r.GET("/", helloWorld)
	// 使用postman在body的form-data中传入对应的参数即可
	_ = r.Run("127.0.0.1:8888")
}
