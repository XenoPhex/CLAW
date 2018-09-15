package exec_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "code.cloudfoundry.org/claw/exec"
)

var _ = Describe("Versions", func() {
	DescribeTable("errors",
		func(value string, expectedErr error) {
			var versions Versions
			err := versions.UnmarshalFlag(value)
			Expect(err).To(MatchError(expectedErr))
			Expect(versions.List).To(BeEmpty())
		},
		Entry("when provided invalid semvers", "fffff,fffff", errors.New("No Major.Minor.Patch elements found")),
	)

	DescribeTable("valid values",
		func(value string, expectedVersions []string) {
			var versions Versions
			err := versions.UnmarshalFlag(value)
			Expect(err).ToNot(HaveOccurred())
			Expect(versions.List).To(Equal(expectedVersions))
		},
		Entry("provided a single valid semvar",
			"1.0.0",
			[]string{"1.0.0"}),
		Entry("provided a sorted list of valid semvers",
			"1.0.0,1.1.1,2.1.3",
			[]string{"1.0.0", "1.1.1", "2.1.3"}),
		Entry("sorts an unsorted list of valid semvers",
			"1.0.0,2.1.3,1.1.1",
			[]string{"1.0.0", "1.1.1", "2.1.3"}),
	)
})
