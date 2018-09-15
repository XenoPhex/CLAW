package request_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"

	. "code.cloudfoundry.org/claw/request"
)

var _ = Describe("Edge", func() {
	var (
		request  *http.Request
		response *http.Response
	)

	BeforeEach(func() {
		var err error
		request, err = http.NewRequest("GET", TestURL, nil)
		Expect(err).ToNot(HaveOccurred())
	})

	Describe("invalid architectures", func() {
		Describe("empty architecture", func() {
			It("returns back a Precondition Failed (412) with a message listing the valid arches", func() {
				response = RunRequest(request, Edge)
				Expect(response.StatusCode).To(Equal(http.StatusPreconditionFailed))
				Eventually(BufferReader(response.Body)).Should(Say("Invalid 'arch' value, please select one of the following architectures: linux32, linux64, macosx64, windows32, windows64"))
			})
		})

		Describe("unsupported architecture", func() {
			BeforeEach(func() {
				AddQuery(request, "arch", "solaris")
			})

			It("returns back a Precondition Failed (412) with a message listing the valid arches", func() {
				response = RunRequest(request, Edge)
				Expect(response.StatusCode).To(Equal(http.StatusPreconditionFailed))
				Eventually(BufferReader(response.Body)).Should(Say("Invalid 'arch' value, please select one of the following architectures: linux32, linux64, macosx64, windows32, windows64"))
			})
		})
	})

	DescribeTable("valid architectures",
		func(arch string, expectedFilename string) {
			AddQuery(request, "arch", arch)
			response = RunRequest(request, Edge)
			Expect(response.StatusCode).To(Equal(http.StatusFound))
			Expect(response.Header.Get("Location")).To(MatchRegexp("https://cf-cli-releases.s3.amazonaws.com/master/%s", expectedFilename))
		},
		Entry("32 bit Linux binary", "linux32", "cf-cli_edge_linux_i686.tgz"),
		Entry("64 bit Linux binary", "linux64", "cf-cli_edge_linux_x86-64.tgz"),
		Entry("64 bit Mac OS X binary", "macosx64", "cf-cli_edge_osx.tgz"),
		Entry("32 bit Windows binary", "windows32", "cf-cli_edge_win32.zip"),
		Entry("64 bit Windows binary", "windows64", "cf-cli_edge_winx64.zip"),
	)
})
