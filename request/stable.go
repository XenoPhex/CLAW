package request

import (
	"net/http"

	"code.cloudfoundry.org/claw/request/internal"
	"github.com/gin-gonic/gin"
)

var StableVersions []string

func Stable(c *gin.Context) {
	requestedArch, _ := c.GetQuery("release")
	version, providedVersion := c.GetQuery("version")
	if !providedVersion {
		version = StableVersions[len(StableVersions)-1]
	} else if invalidVersion(version) {
		internal.InvalidReleaseVersionError(StableVersions, c)
		return
	}

	if !internal.SupportedStableArch(requestedArch) {
		internal.InvalidArchError("release", internal.StableArches(), c)
		return
	}
	c.Redirect(http.StatusFound, internal.StableURL(requestedArch, version))
}

func invalidVersion(version string) bool {
	for _, releaseVersion := range StableVersions {
		if releaseVersion == version {
			return false
		}
	}
	return true
}
