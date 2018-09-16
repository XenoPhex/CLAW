package request

import (
	"github.com/gin-gonic/gin"
)

func Homebrew(c *gin.Context) {
	version := extractVersion(c.Param("filename"))
	redirectToStable("macosx64-binary", version, c)
}

// extractVersion will take a filename in the format of 'cf-x.x.x.tgz' and
// return back 'x.x.x'.
func extractVersion(filename string) string {
	if len(filename) < 9 {
		return ""
	}
	return filename[:len(filename)-4][4:]
}
