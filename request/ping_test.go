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
		router.GET("/ping", Ping)

		var err error
		request, err = http.NewRequest("GET", "/ping", nil)
		Expect(err).ToNot(HaveOccurred())
	})

	It("returns back Status 200", func() {
		response = RunRequest(request)
		Expect(response.StatusCode).To(Equal(http.StatusOK))
	})

	It("returns back pong", func() {
		response = RunRequest(request)
		Eventually(BufferReader(response.Body)).Should(Say("pong"))
	})
})
