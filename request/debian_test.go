package request_test

import (
	"fmt"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "code.cloudfoundry.org/claw/request"
)

var _ = Describe("Debian Repository", func() {
	Describe("DebianDist", func() {
		BeforeEach(func() {
			router.GET("/debian/dists/*page", DebianDist)
		})

		It("redirects to the equivilant /debian/dists URL", func() {
			request, err := http.NewRequest("GET", "/debian/dists/bob", nil)
			Expect(err).ToNot(HaveOccurred())

			response := RunRequest(request)
			Expect(response.StatusCode).To(Equal(http.StatusFound))
			Expect(response.Header.Get("Location")).To(Equal("https://cf-cli-debian-repo.s3.amazonaws.com/dists/bob"))
		})
	})

	Describe("DebianPool", func() {
		BeforeEach(func() {
			StableVersions = []string{"6.0.0", "6.1.0", "6.1.1", "6.2.0"}

			router.GET("/debian/pool/*page", DebianPool)
		})

		Describe("a non-versioned URI", func() {
			// Note: Unable to find an example of this, writing this test based off
			// of code from original version of CLAW. Who knows what this link
			// actually looks like.
			DescribeTable("without version redirects to latested binary for specified architecture",
				func(arch string, expectedFilename string) {
					Skip("Don't implement this until someone complains")
					request, err := http.NewRequest("GET",
						fmt.Sprintf("/debian/pool/stable/c/cf/????=%s", arch),
						nil)
					Expect(err).ToNot(HaveOccurred())
					response := RunRequest(request)

					Expect(response.StatusCode).To(Equal(http.StatusFound))
					requestedVersion := StableVersions[len(StableVersions)-1]
					expectedFilename = fmt.Sprintf(expectedFilename, requestedVersion)
					Expect(response.Header.Get("Location")).To(MatchRegexp("https://s3-us-west-1.amazonaws.com/cf-cli-releases/releases/v%s/%s", requestedVersion, expectedFilename))
				},
				Entry("32 bit Debian installer", "i686", "cf-cli-installer_%s_i686.deb"),
				Entry("64 bit Debian installer", "x86-64", "cf-cli-installer_%s_x86-64.deb"),
			)
		})

		Describe("a versioned URI", func() {
			DescribeTable("specified architecture/version redirects to specified binary",
				func(arch string, expectedFilename string) {
					requestedVersion := "6.1.0"
					request, err := http.NewRequest("GET",
						fmt.Sprintf("/debian/pool/stable/c/cf/cf-cli-installer_%s_%s.deb", requestedVersion, arch),
						nil)
					Expect(err).ToNot(HaveOccurred())
					response := RunRequest(request)

					Expect(response.StatusCode).To(Equal(http.StatusFound))
					expectedFilename = fmt.Sprintf(expectedFilename, requestedVersion)
					Expect(response.Header.Get("Location")).To(MatchRegexp("https://s3-us-west-1.amazonaws.com/cf-cli-releases/releases/v%s/%s", requestedVersion, expectedFilename))
				},
				Entry("32 bit Debian installer", "i686", "cf-cli-installer_%s_i686.deb"),
				Entry("64 bit Debian installer", "x86-64", "cf-cli-installer_%s_x86-64.deb"),
			)
		})
	})
})
