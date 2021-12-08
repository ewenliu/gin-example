package cat

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func barkHandler(c *gin.Context) {
	c.String(http.StatusOK, "this is cat bark handler")
}

func otherHandler(c *gin.Context) {
	c.String(http.StatusOK, "this is cat other handler")
}
