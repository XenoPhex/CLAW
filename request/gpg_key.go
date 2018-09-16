package request

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var GPGKeyBody string

func GPGKey(c *gin.Context) {
	c.String(http.StatusOK, GPGKeyBody)
}
