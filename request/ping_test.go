package request_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"

	. "code.cloudfoundry.org/claw/request"
)

var _ = Describe("Ping", func() {
	var (
		request  *http.Request
		response *http.Response
	)

	BeforeEach(func() {
		var err error
		request, err = http.NewRequest("GET", TestURL, nil)
		Expect(err).ToNot(HaveOccurred())
	})

	JustBeforeEach(func() {
		response = RunRequest(request, Ping)
	})

	It("returns back Status 200", func() {
		Expect(response.StatusCode).To(Equal(http.StatusOK))
	})

	It("returns back pong", func() {
		Eventually(BufferReader(response.Body)).Should(Say("pong"))
	})
})
