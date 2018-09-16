package internal

import (
	"fmt"
	"sort"
)

var stableFilenames = map[string]string{
	"debian32":        "cf-cli-installer_%s_i686.deb",
	"debian64":        "cf-cli-installer_%s_x86-64.deb",
	"redhat32":        "cf-cli-installer_%s_i686.rpm",
	"redhat64":        "cf-cli-installer_%s_x86-64.rpm",
	"macosx64":        "cf-cli-installer_%s_osx.pkg",
	"windows32":       "cf-cli-installer_%s_win32.zip",
	"windows64":       "cf-cli-installer_%s_winx64.zip",
	"linux32-binary":  "cf-cli_%s_linux_i686.tgz",
	"linux64-binary":  "cf-cli_%s_linux_x86-64.tgz",
	"macosx64-binary": "cf-cli_%s_osx.tgz",
	"windows32-exe":   "cf-cli_%s_win32.zip",
	"windows64-exe":   "cf-cli_%s_winx64.zip",
}

func StableArches() []string {
	var arches []string
	for supported := range stableFilenames {
		arches = append(arches, supported)
	}

	list := sort.StringSlice(arches)
	list.Sort()
	return list
}

func StableURL(arch string, version string) string {
	filename := fmt.Sprintf(stableFilenames[arch], version)
	return fmt.Sprintf("https://s3-us-west-1.amazonaws.com/cf-cli-releases/releases/v%s/%s", version, filename)
}

func StableURLFromFile(filename string, version string) string {
	return fmt.Sprintf("https://s3-us-west-1.amazonaws.com/cf-cli-releases/releases/v%s/%s", version, filename)
}

func SupportedStableArch(arch string) bool {
	for supported := range stableFilenames {
		if arch == supported {
			return true
		}
	}
	return false
}
