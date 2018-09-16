package request_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"

	. "code.cloudfoundry.org/claw/request"
)

var _ = Describe("GPG Key", func() {
	var (
		request  *http.Request
		response *http.Response
	)

	BeforeEach(func() {
		var err error
		request, err = http.NewRequest("GET", TestURL, nil)
		Expect(err).ToNot(HaveOccurred())
		GPGKeyBody = "some-GPG-key"
	})

	It("returns the GPG key", func() {
		response = RunRequest(request, GPGKey)
		Expect(response.StatusCode).To(Equal(http.StatusOK))
		Expect(response.Header.Get("Content-Type")).To(Equal("text/plain; charset=utf-8"))
		Eventually(BufferReader(response.Body)).Should(Say("some-GPG-key"))
	})
})
