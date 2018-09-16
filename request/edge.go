package request

import (
	"net/http"

	"code.cloudfoundry.org/claw/request/internal"
	"github.com/gin-gonic/gin"
)

func Edge(c *gin.Context) {
	requestedArch := c.Query("arch")
	if !internal.SupportedEdgeArch(requestedArch) {
		internal.InvalidArchError("arch", internal.EdgeArches(), c)
		return
	}
	c.Redirect(http.StatusFound, internal.EdgeURL(requestedArch))
}
