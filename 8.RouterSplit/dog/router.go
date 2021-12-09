package dog

import "github.com/gin-gonic/gin"

func LoadHandler(e *gin.Engine) {
	routerGroup := e.Group("/dog")

	{
		routerGroup.GET("/bark", barkHandler)
		routerGroup.POST("/other", otherHandler)
	}

	deepRouterGroup := routerGroup.Group("/info")

	{
		deepRouterGroup.GET("/overview", barkHandler)
	}
}
