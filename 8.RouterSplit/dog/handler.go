package dog

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func barkHandler(c *gin.Context) {
	c.String(http.StatusOK, "this is dog bark handler")
}

func otherHandler(c *gin.Context) {
	c.String(http.StatusOK, "this is dog other handler")
}

