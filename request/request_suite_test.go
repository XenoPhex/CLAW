package request_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	// TestURL is the URL that all the requests are tested as. Since the routing
	// is not being tested, this could be anything.
	TestURL = "/some-url"
)

var router *gin.Engine

func TestRequest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Request Suite")
}

var _ = SynchronizedBeforeSuite(func() []byte {
	return nil
}, func(data []byte) {
	gin.DefaultWriter = &bytes.Buffer{}      // disable gin's logging
	gin.DefaultErrorWriter = &bytes.Buffer{} // disable gin's logging
	gin.SetMode(gin.TestMode)                // disable gin's logging
})

var _ = BeforeEach(func() {
	router = gin.Default()
})

func AddQuery(request *http.Request, name string, value string) {
	queries := request.URL.Query()
	queries.Add(name, value)
	request.URL.RawQuery = queries.Encode()
}

func RunRequest(request *http.Request) *http.Response {
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)
	return w.Result()
}
