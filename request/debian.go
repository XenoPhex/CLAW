package request

import (
	"fmt"
	"net/http"
	"path/filepath"
	"regexp"

	"code.cloudfoundry.org/claw/request/internal"
	"github.com/gin-gonic/gin"
)

const debianRepoRoot = "https://cf-cli-debian-repo.s3.amazonaws.com"

var versionRegexp = regexp.MustCompile(".*installer_(.*)_.*")

// DebianDist redirects URLs for all the repository files related to the
// 'apt' command. These do not include the *.deb binary URLs.
func DebianDist(c *gin.Context) {
	c.Redirect(http.StatusFound, fmt.Sprintf("%s/dists%s", debianRepoRoot, c.Param("page")))
}

// DebianPool redirects URLs for the *.deb binary downloads to the standard
// stable link.
func DebianPool(c *gin.Context) {
	filename := filepath.Base(c.Param("page"))
	if matches := versionRegexp.FindAllStringSubmatch(filename, -1); len(matches) > 0 {
		version := matches[0][1]
		c.Redirect(http.StatusFound, internal.StableURLFromFile(filename, version))
		return
	}

	// TODO: The previous version of CLAW would redirect to the latest version of
	// the deb installer. Since there have been no recent records/logs of this
	// extra behavior, skipping implementation for now.
	c.String(http.StatusNotImplemented, "If you are seeing this message please file a bug report on github.com/cloudfoundry/claw")
}
