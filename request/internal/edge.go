package internal

import (
	"fmt"
	"sort"
)

var edgeFilenames = map[string]string{
	"linux32":   "cf-cli_edge_linux_i686.tgz",
	"linux64":   "cf-cli_edge_linux_x86-64.tgz",
	"macosx64":  "cf-cli_edge_osx.tgz",
	"windows32": "cf-cli_edge_win32.zip",
	"windows64": "cf-cli_edge_winx64.zip",
}

func EdgeArches() []string {
	var arches []string
	for supported := range edgeFilenames {
		arches = append(arches, supported)
	}

	list := sort.StringSlice(arches)
	list.Sort()
	return list
}

func EdgeURL(arch string) string {
	return fmt.Sprintf("https://cf-cli-releases.s3.amazonaws.com/master/%s", edgeFilenames[arch])
}

func SupportedEdgeArch(arch string) bool {
	for supported := range edgeFilenames {
		if arch == supported {
			return true
		}
	}
	return false
}
