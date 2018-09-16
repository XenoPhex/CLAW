package request_test

import (
	"fmt"
	"net/http"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"

	. "code.cloudfoundry.org/claw/request"
)

var _ = Describe("Stable", func() {
	var (
		request  *http.Request
		response *http.Response
	)

	BeforeEach(func() {
		StableVersions = []string{"6.0.0", "6.1.0", "6.1.1", "6.2.0"}

		router.GET("/stable", Stable)

		var err error
		request, err = http.NewRequest("GET", "/stable", nil)
		Expect(err).ToNot(HaveOccurred())
	})

	Describe("invalid query parameters", func() {
		Describe("empty architecture", func() {
			It("returns back a Precondition Failed (412) with a message listing the valid arches", func() {
				response = RunRequest(request)
				Expect(response.StatusCode).To(Equal(http.StatusPreconditionFailed))
				Eventually(BufferReader(response.Body)).Should(Say("Invalid 'release' value, please select one of the following architectures: debian32, debian64, linux32-binary, linux64-binary, macosx64, macosx64-binary, redhat32, redhat64, windows32, windows32-exe, windows64, windows64-exe"))
			})
		})

		Describe("unsupported architecture", func() {
			BeforeEach(func() {
				AddQuery(request, "release", "solaris")
			})

			It("returns back a Precondition Failed (412) with a message listing the valid arches", func() {
				response = RunRequest(request)
				Expect(response.StatusCode).To(Equal(http.StatusPreconditionFailed))
				Eventually(BufferReader(response.Body)).Should(Say("Invalid 'release' value, please select one of the following architectures: debian32, debian64, linux32-binary, linux64-binary, macosx64, macosx64-binary, redhat32, redhat64, windows32, windows32-exe, windows64, windows64-exe"))
			})
		})

		Describe("invalid version number", func() {
			BeforeEach(func() {
				AddQuery(request, "version", "banana")
			})

			It("returns back a Precondition Failed (412) with a message listing the valid releases", func() {
				response = RunRequest(request)
				Expect(response.StatusCode).To(Equal(http.StatusPreconditionFailed))
				Eventually(BufferReader(response.Body)).Should(Say("Invalid 'version' value, please select one of the following versions: %s", strings.Join(StableVersions, ", ")))
			})
		})
	})

	Describe("valid query parameters", func() {
		DescribeTable("specified architecture/version redirects to specified binary",
			func(arch string, expectedFilename string) {
				requestedVersion := "6.1.0"
				AddQuery(request, "release", arch)
				AddQuery(request, "version", requestedVersion)
				response = RunRequest(request)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				expectedFilename = fmt.Sprintf(expectedFilename, requestedVersion)
				Expect(response.Header.Get("Location")).To(MatchRegexp("https://s3-us-west-1.amazonaws.com/cf-cli-releases/releases/v%s/%s", requestedVersion, expectedFilename))
			},
			Entry("32 bit Debian installer", "debian32", "cf-cli-installer_%s_i686.deb"),
			Entry("64 bit Debian installer", "debian64", "cf-cli-installer_%s_x86-64.deb"),
			Entry("32 bit Redhat installer", "redhat32", "cf-cli-installer_%s_i686.rpm"),
			Entry("64 bit Redhat installer", "redhat64", "cf-cli-installer_%s_x86-64.rpm"),
			Entry("64 bit Mac OS X installer", "macosx64", "cf-cli-installer_%s_osx.pkg"),
			Entry("32 bit Windows installer", "windows32", "cf-cli-installer_%s_win32.zip"),
			Entry("64 bit Windows installer", "windows64", "cf-cli-installer_%s_winx64.zip"),
			Entry("32 bit Linux binary", "linux32-binary", "cf-cli_%s_linux_i686.tgz"),
			Entry("64 bit Linux binary", "linux64-binary", "cf-cli_%s_linux_x86-64.tgz"),
			Entry("64 bit Mac OS X binary", "macosx64-binary", "cf-cli_%s_osx.tgz"),
			Entry("32 bit Windows binary", "windows32-exe", "cf-cli_%s_win32.zip"),
			Entry("64 bit Windows binary", "windows64-exe", "cf-cli_%s_winx64.zip"),
		)

		DescribeTable("valid architectures without a specified version for latest",
			func(arch string, expectedFilename string) {
				AddQuery(request, "release", arch)
				response = RunRequest(request)

				Expect(response.StatusCode).To(Equal(http.StatusFound))

				requestedVersion := "6.2.0"
				expectedFilename = fmt.Sprintf(expectedFilename, requestedVersion)
				Expect(response.Header.Get("Location")).To(MatchRegexp("https://s3-us-west-1.amazonaws.com/cf-cli-releases/releases/v%s/%s", requestedVersion, expectedFilename))
			},
			Entry("32 bit Debian installer", "debian32", "cf-cli-installer_%s_i686.deb"),
			Entry("64 bit Debian installer", "debian64", "cf-cli-installer_%s_x86-64.deb"),
			Entry("32 bit Redhat installer", "redhat32", "cf-cli-installer_%s_i686.rpm"),
			Entry("64 bit Redhat installer", "redhat64", "cf-cli-installer_%s_x86-64.rpm"),
			Entry("64 bit Mac OS X installer", "macosx64", "cf-cli-installer_%s_osx.pkg"),
			Entry("32 bit Windows installer", "windows32", "cf-cli-installer_%s_win32.zip"),
			Entry("64 bit Windows installer", "windows64", "cf-cli-installer_%s_winx64.zip"),
			Entry("32 bit Linux binary", "linux32-binary", "cf-cli_%s_linux_i686.tgz"),
			Entry("64 bit Linux binary", "linux64-binary", "cf-cli_%s_linux_x86-64.tgz"),
			Entry("64 bit Mac OS X binary", "macosx64-binary", "cf-cli_%s_osx.tgz"),
			Entry("32 bit Windows binary", "windows32-exe", "cf-cli_%s_win32.zip"),
			Entry("64 bit Windows binary", "windows64-exe", "cf-cli_%s_winx64.zip"),
		)
	})
})
