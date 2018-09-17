package request_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "code.cloudfoundry.org/claw/request"
)

var _ = Describe("Fedora/Redhat Repository", func() {
	Describe("FedoraRepo", func() {
		BeforeEach(func() {
			router.GET("/fedora/cloudfoundry-cli.repo", FedoraRepo)
		})

		It("redirects to the Fedora repository file", func() {
			request, err := http.NewRequest("GET", "/fedora/cloudfoundry-cli.repo", nil)
			Expect(err).ToNot(HaveOccurred())

			response := RunRequest(request)
			Expect(response.StatusCode).To(Equal(http.StatusFound))
			Expect(response.Header.Get("Location")).To(Equal("https://cf-cli-rpm-repo.s3.amazonaws.com/cloudfoundry-cli.repo"))
		})
	})

	Describe("FedoraRepoData", func() {
		BeforeEach(func() {
			router.GET("/fedora/repodata/*page", FedoraRepoData)
		})

		It("redirects to the equivilant /fedora/repodata URL", func() {
			request, err := http.NewRequest("GET", "/fedora/repodata/bob", nil)
			Expect(err).ToNot(HaveOccurred())

			response := RunRequest(request)
			Expect(response.StatusCode).To(Equal(http.StatusFound))
			Expect(response.Header.Get("Location")).To(Equal("https://cf-cli-rpm-repo.s3.amazonaws.com/bob"))
		})
	})
})
