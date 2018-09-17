package request

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const fedoraRepoRoot = "https://cf-cli-rpm-repo.s3.amazonaws.com"

func FedoraRepo(c *gin.Context) {
	c.Redirect(http.StatusFound, fmt.Sprintf("%s/cloudfoundry-cli.repo", fedoraRepoRoot))
}
