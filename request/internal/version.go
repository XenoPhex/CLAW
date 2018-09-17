package internal

import "regexp"

var VersionRegexp = regexp.MustCompile(".*installer_(.*)_.*")
