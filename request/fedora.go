package request

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

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

// FedoraReleases redirects to the fedora binary RPMs for that
// architecture.
func FedoraReleases(c *gin.Context) {
	filename := filepath.Base(c.Param("page"))
	filename = strings.Replace(filename, "_linux", "", -1)
	if matches := internal.VersionRegexp.FindAllStringSubmatch(filename, -1); len(matches) > 0 {
		version := matches[0][1]
		c.Redirect(http.StatusFound, internal.StableURLFromFile(filename, version))
		return
	}

	internal.InvalidReleaseVersionError(StableVersions, c)
}
