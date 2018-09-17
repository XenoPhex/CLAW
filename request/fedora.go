package request

import (
	"fmt"
	"net/http"

	"code.cloudfoundry.org/claw/request/internal"
	"github.com/gin-gonic/gin"
)

// FedoraRepoData redirect to the repository metadata files. These do not
// include the RPM binaries.
func FedoraRepoData(c *gin.Context) {
	c.Redirect(http.StatusFound, internal.FedoraRepoURL(fmt.Sprintf("/repodata%s", c.Param("page"))))
}

// FedoraUserConfig provides the a config file that can be used to access the
// repository on the local machine.
func FedoraUserConfig(c *gin.Context) {
	c.Redirect(http.StatusFound, internal.FedoraRepoURL("/cloudfoundry-cli.repo"))
}
