package request

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const debianRepoRoot = "https://cf-cli-debian-repo.s3.amazonaws.com"

func DebianDist(c *gin.Context) {
	c.Redirect(http.StatusFound, fmt.Sprintf("%s/dists%s", debianRepoRoot, c.Param("page")))
}
