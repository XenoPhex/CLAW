package internal

import "fmt"

const fedoraRepoRoot = "https://cf-cli-rpm-repo.s3.amazonaws.com"

func FedoraRepoURL(uri string) string {
	return fmt.Sprintf("%s%s", fedoraRepoRoot, uri)
}
