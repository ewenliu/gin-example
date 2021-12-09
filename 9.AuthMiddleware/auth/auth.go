package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var Keys map[interface{}]interface{}

func LoginAuth() gin.HandlerFunc{
	return func(c *gin.Context) {
		_, exist := Keys["cookie"]
		if !exist{
			c.Abort()
			c.String(http.StatusUnauthorized, "Login required!")
		}
		c.Next()
	}
}

func SetKeys(key, val interface{}){
	Keys[key] = val
}

func init()  {
	Keys = make(map[interface{}]interface{})
}
