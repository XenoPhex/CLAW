package request_test

import (
	"net/http"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"

	. "code.cloudfoundry.org/claw/request"
)

var _ = Describe("Fedora/Redhat Repository", func() {
	Describe("FedoraRepoData", func() {
		BeforeEach(func() {
			router.GET("/fedora/repodata/*page", FedoraRepoData)
		})

		It("redirects to the equivilant /fedora/repodata URL", func() {
			request, err := http.NewRequest("GET", "/fedora/repodata/bob", nil)
			Expect(err).ToNot(HaveOccurred())

			response := RunRequest(request)
			Expect(response.StatusCode).To(Equal(http.StatusFound))
			Expect(response.Header.Get("Location")).To(Equal("https://cf-cli-rpm-repo.s3.amazonaws.com/repodata/bob"))
		})
	})

	Describe("FedoraUserConfig", func() {
		BeforeEach(func() {
			router.GET("/fedora/cloudfoundry-cli.repo", FedoraUserConfig)
		})

		It("redirects to the Fedora repository configuration file", func() {
			request, err := http.NewRequest("GET", "/fedora/cloudfoundry-cli.repo", nil)
			Expect(err).ToNot(HaveOccurred())

			response := RunRequest(request)
			Expect(response.StatusCode).To(Equal(http.StatusFound))
			Expect(response.Header.Get("Location")).To(Equal("https://cf-cli-rpm-repo.s3.amazonaws.com/cloudfoundry-cli.repo"))
		})
	})

	Describe("FedoraReleases", func() {
		BeforeEach(func() {
			StableVersions = []string{"6.0.0", "6.1.0", "6.1.1", "6.2.0"}

			router.GET("/fedora/releases/*page", FedoraReleases)
		})

		DescribeTable("specified architecture/version redirects to specified binary",
			func(requestedURL string, expectedURL string) {
				request, err := http.NewRequest("GET", requestedURL, nil)
				Expect(err).ToNot(HaveOccurred())
				response := RunRequest(request)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(response.Header.Get("Location")).To(Equal(expectedURL))
			},

			Entry("32 bit Fedora installer",
				"/fedora/releases/v6.39.0/cf-cli-installer_6.39.0_i686.rpm",
				"https://s3-us-west-1.amazonaws.com/cf-cli-releases/releases/v6.39.0/cf-cli-installer_6.39.0_i686.rpm"),
			Entry("64 bit Fedora installer",
				"/fedora/releases/v6.39.0/cf-cli-installer_6.39.0_x86-64.rpm",
				"https://s3-us-west-1.amazonaws.com/cf-cli-releases/releases/v6.39.0/cf-cli-installer_6.39.0_x86-64.rpm"),

			Entry("32 bit Fedora installer with linux in the name",
				"/fedora/releases/v6.12.3/cf-cli-installer_6.12.3_linux_i686.rpm",
				"https://s3-us-west-1.amazonaws.com/cf-cli-releases/releases/v6.12.3/cf-cli-installer_6.12.3_i686.rpm"),
			Entry("64 bit Fedora installer with linux in the name",
				"/fedora/releases/v6.12.3/cf-cli-installer_6.12.3_linux_x86-64.rpm",
				"https://s3-us-west-1.amazonaws.com/cf-cli-releases/releases/v6.12.3/cf-cli-installer_6.12.3_x86-64.rpm"),
		)

		When("version is not specified", func() {
			It("returns an invalid version error", func() {
				request, err := http.NewRequest("GET", "/fedora/releases/potato.rpm", nil)
				Expect(err).ToNot(HaveOccurred())
				response := RunRequest(request)

				Expect(response.StatusCode).To(Equal(http.StatusPreconditionFailed))
				Eventually(BufferReader(response.Body)).Should(Say("Invalid 'version' value, please select one of the following versions: %s", strings.Join(StableVersions, ", ")))
			})
		})
	})
})
