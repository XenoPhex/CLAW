package request

import (
	"fmt"
	"net/http"

	"code.cloudfoundry.org/claw/request/internal"
	"github.com/gin-gonic/gin"
)

func FedoraRepo(c *gin.Context) {
	c.Redirect(http.StatusFound, internal.FedoraRepoURL("/cloudfoundry-cli.repo"))
}

func FedoraRepoData(c *gin.Context) {
	c.Redirect(http.StatusFound, internal.FedoraRepoURL(fmt.Sprintf("/repodata%s", c.Param("page"))))
}
