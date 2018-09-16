package request_test

import (
	"fmt"
	"net/http"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"

	. "code.cloudfoundry.org/claw/request"
)

var _ = Describe("homebrew", func() {
	var (
		request  *http.Request
		response *http.Response
	)

	BeforeEach(func() {
		StableVersions = []string{"6.0.0", "6.1.0", "6.1.1", "6.2.0"}

		router.GET("/homebrew/*filename", Homebrew)
	})

	Describe("invalid url parameters", func() {
		var testURL string

		JustBeforeEach(func() {
			var err error
			request, err = http.NewRequest("GET", testURL, nil)
			Expect(err).ToNot(HaveOccurred())
		})

		Describe("empty filename", func() {
			BeforeEach(func() {
				testURL = "/homebrew/"
			})

			It("returns back a Precondition Failed (412) with a message listing the valid releases", func() {
				response = RunRequest(request)
				Expect(response.StatusCode).To(Equal(http.StatusPreconditionFailed))
				Eventually(BufferReader(response.Body)).Should(Say("Invalid 'version' value, please select one of the following versions: %s", strings.Join(StableVersions, ", ")))
			})
		})

		Describe("unsupported version", func() {
			BeforeEach(func() {
				testURL = "/homebrew/banana.tgz"
			})

			It("returns back a Precondition Failed (412) with a message listing the valid releases", func() {
				response = RunRequest(request)
				Expect(response.StatusCode).To(Equal(http.StatusPreconditionFailed))
				Eventually(BufferReader(response.Body)).Should(Say("Invalid 'version' value, please select one of the following versions: %s", strings.Join(StableVersions, ", ")))
			})
		})
	})

	Describe("valid url parameters", func() {
		It("redirects to the Mac OS X binary for that given version", func() {
			for _, requestedVersion := range StableVersions {
				request, err := http.NewRequest("GET", fmt.Sprintf("/homebrew/cf-%s.tgz", requestedVersion), nil)
				Expect(err).ToNot(HaveOccurred())

				response = RunRequest(request)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(response.Header.Get("Location")).To(MatchRegexp("https://s3-us-west-1.amazonaws.com/cf-cli-releases/releases/v%s/cf-cli_%s_osx.tgz", requestedVersion, requestedVersion))
			}
		})
	})
})
