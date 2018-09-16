package request_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	// . "github.com/onsi/gomega/gbytes"

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
})
