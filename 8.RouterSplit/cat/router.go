package cat

import "github.com/gin-gonic/gin"

func LoadHandler(e *gin.Engine) {
	routerGroup := e.Group("/cat")

	{
		routerGroup.GET("/bark", barkHandler)
		routerGroup.POST("/other", otherHandler)
	}
}
