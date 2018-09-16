package exec_test

import (
	"os"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "code.cloudfoundry.org/claw/exec"
	"code.cloudfoundry.org/claw/request"
)

func serverArgsMinus(argsToSkip ...string) []string {
	testArgs := map[string]string{
		"--available-versions": "1.0.0,1.1.0,1.1.1,1.2.0",
		"--gpg-key":            "some-gpg-key",
		"--port":               "8081",
		"--release":            "",
	}

	for _, arg := range argsToSkip {
		delete(testArgs, arg)
	}

	var serverArgs []string
	for key, val := range testArgs {
		serverArgs = append(serverArgs, key, val)
	}
	return serverArgs
}

var _ = Describe("Setup", func() {
	Describe("Available Stable V6 Versions", func() {
		When("not provided", func() {
			It("returns an error", func() {
				_, err := Setup(serverArgsMinus("--available-versions"))
				Expect(err).To(MatchError("the required flag `--available-versions' was not specified"))
			})
		})

		When("provided via flag", func() {
			It("sets the stable versions to the provided value", func() {
				_, err := Setup(serverArgsMinus())
				Expect(err).ToNot(HaveOccurred())
				Expect(request.StableVersions).To(ConsistOf("1.0.0", "1.1.0", "1.1.1", "1.2.0"))
			})
		})

		When("provided via Environment Variable", func() {
			BeforeEach(func() {
				os.Setenv("AVAILABLE_VERSIONS", "1.0.0,1.1.0,1.1.1,1.2.0")
			})

			AfterEach(func() {
				os.Unsetenv("AVAILABLE_VERSIONS")
			})

			It("sets the stable versions to the provided value", func() {
				_, err := Setup(serverArgsMinus("--available-versions"))
				Expect(err).ToNot(HaveOccurred())
				Expect(request.StableVersions).To(ConsistOf("1.0.0", "1.1.0", "1.1.1", "1.2.0"))
			})
		})
	})

	Describe("GPG Key", func() {
		When("not provided", func() {
			It("returns an error", func() {
				_, err := Setup(serverArgsMinus("--gpg-key"))
				Expect(err).To(MatchError("the required flag `--gpg-key' was not specified"))
			})
		})

		When("provided via flag", func() {
			It("sets the gpg key to the provided value", func() {
				_, err := Setup(serverArgsMinus())
				Expect(err).ToNot(HaveOccurred())
				Expect(request.GPGKeyBody).To(Equal("some-gpg-key"))
			})
		})

		When("provided via Environment Variable", func() {
			BeforeEach(func() {
				os.Setenv("GPG_KEY", "some-gpg-key")
			})

			AfterEach(func() {
				os.Unsetenv("GPG_KEY")
			})

			It("sets the gpg key to the provided value", func() {
				_, err := Setup(serverArgsMinus("--gpg-key"))
				Expect(err).ToNot(HaveOccurred())
				Expect(request.GPGKeyBody).To(Equal("some-gpg-key"))
			})
		})
	})

	Describe("Port", func() {
		When("not provided", func() {
			It("defaults to 8080", func() {
				listenAddr, err := Setup(serverArgsMinus("--port"))
				Expect(err).ToNot(HaveOccurred())
				Expect(listenAddr).To(Equal(":8080"))
			})
		})

		When("provided via flag", func() {
			It("sets the port in the listen address to the provided value", func() {
				listenAddr, err := Setup(serverArgsMinus())
				Expect(err).ToNot(HaveOccurred())
				Expect(listenAddr).To(Equal(":8081"))
			})
		})

		When("provided via Environment Variable", func() {
			BeforeEach(func() {
				os.Setenv("PORT", "8081")
			})

			AfterEach(func() {
				os.Unsetenv("PORT")
			})

			It("sets the port in the listen address to the provided value", func() {
				listenAddr, err := Setup(serverArgsMinus("--port"))
				Expect(err).ToNot(HaveOccurred())
				Expect(listenAddr).To(Equal(":8081"))
			})
		})
	})

	Describe("Release", func() {
		When("not provided", func() {
			It("defaults to debug mode", func() {
				_, err := Setup(serverArgsMinus("--release"))
				Expect(err).ToNot(HaveOccurred())
				Expect(gin.Mode()).To(Equal(gin.DebugMode))
			})
		})

		When("provided via flag", func() {
			It("runs in release mode", func() {
				_, err := Setup(serverArgsMinus())
				Expect(err).ToNot(HaveOccurred())
				Expect(gin.Mode()).To(Equal(gin.ReleaseMode))
			})
		})
	})
})
