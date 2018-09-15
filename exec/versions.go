package exec

import (
	"sort"
	"strings"

	"github.com/blang/semver"
)

type Versions struct {
	List []string
}

func (versions *Versions) UnmarshalFlag(value string) error {
	split := strings.Split(value, ",")

	var validVersions semver.Versions
	for _, version := range split {
		validVersion, err := semver.Make(version)
		if err != nil {
			return err
		}
		validVersions = append(validVersions, validVersion)
	}
	sort.Sort(validVersions)

	for _, version := range validVersions {
		versions.List = append(versions.List, version.String())
	}
	return nil
}
