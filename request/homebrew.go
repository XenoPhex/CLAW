package request

import (
	"errors"
	"net/http"

	"code.cloudfoundry.org/claw/request/internal"
	"github.com/gin-gonic/gin"
)

func Homebrew(c *gin.Context) {
	version, err := extractVersion(c.Param("filename"))
	if err != nil || invalidVersion(version) {
		internal.InvalidReleaseVersionError(StableVersions, c)
		return
	}
	c.Redirect(http.StatusFound, internal.StableURL("macosx64-binary", version))
}

// extractVersion will take a filename in the format of 'cf-x.x.x.tgz' and
// return back 'x.x.x'.
func extractVersion(filename string) (string, error) {
	if len(filename) < 9 {
		return "", errors.New("filename too short")
	}
	version := filename[:len(filename)-4][4:]
	return version, nil
}
