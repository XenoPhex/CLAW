package exec

import (
	"fmt"

	"code.cloudfoundry.org/claw/request"
	"github.com/gin-gonic/gin"
	flags "github.com/jessevdk/go-flags"
)

type Opts struct {
	AvailableVersions Versions `long:"available-versions" required:"true" value-name:"0.0.0[,0.1.0...]" env:"AVAILABLE_VERSIONS" description:"List of available stable V6 versions"`
	GPGKey            string   `long:"gpg-key" required:"true" env:"GPG_KEY" description:"Public GPG Key used for the Debian and Redhat repositories"`
	Port              int      `long:"port" default:"8080" env:"PORT" description:"App server port"`
	Release           bool     `long:"release" description:"Enable release mode (aka production) for app server"`
}

func Setup(args []string) (string, error) {
	var opts Opts
	parser := flags.NewParser(&opts, flags.HelpFlag|flags.PassDoubleDash)
	_, err := parser.ParseArgs(args)
	if err != nil {
		return "", err
	}

	request.StableVersions = opts.AvailableVersions.List
	request.GPGKeyBody = opts.GPGKey

	if opts.Release {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	return fmt.Sprintf(":%d", opts.Port), nil
}

func Server(args []string) (*gin.Engine, string, error) {
	listenAddr, err := Setup(args)
	if err != nil {
		return nil, "", err
	}

	r := gin.Default()
	r.GET("/debian/cli.cloudfoundry.org.key", request.GPGKey)
	r.GET("/edge", request.Edge)
	r.GET("/fedora/cli.cloudfoundry.org.key", request.GPGKey)
	r.GET("/ping", request.Ping)
	r.GET("/stable", request.Stable)
	return r, listenAddr, nil
}
