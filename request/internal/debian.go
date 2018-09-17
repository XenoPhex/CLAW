package internal

import "fmt"

const debianRepoRoot = "https://cf-cli-debian-repo.s3.amazonaws.com"

func DebianDistURL(uri string) string {
	return fmt.Sprintf("%s/dists%s", debianRepoRoot, uri)
}
