package cat

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func barkHandler(c *gin.Context) {
	c.String(http.StatusOK, "this is cat bark handler")
}

func infoHandler(c *gin.Context){
	id := c.Param("id")
	fmt.Println(id)
}

func otherHandler(c *gin.Context) {
	c.String(http.StatusOK, "this is cat other handler")
}
