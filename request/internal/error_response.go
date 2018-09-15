package internal

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func InvalidArchError(queryName string, arches []string, c *gin.Context) {
	c.String(http.StatusPreconditionFailed,
		fmt.Sprintf("Invalid '%s' value, please select one of the following architectures: %s",
			queryName,
			strings.Join(arches, ", "),
		),
	)
}

func InvalidReleaseVersionError(versions []string, c *gin.Context) {
	c.String(http.StatusPreconditionFailed,
		fmt.Sprintf("Invalid 'version' value, please select one of the following versions: %s",
			strings.Join(versions, ", "),
		),
	)
}
