package exec

import (
	"fmt"
	"os"

	"code.cloudfoundry.org/claw/request"
	"github.com/gin-gonic/gin"
	flags "github.com/jessevdk/go-flags"
)

type Opts struct {
	AvailableVersions Versions `long:"available-versions" required:"true" value-name:"0.0.0[,0.1.0...]" env:"AVAILABLE_VERSIONS" description:"List of available stable V6 versions"`
	Port              int      `long:"port" default:"8080" env:"PORT" description:"App server port"`
	Release           bool     `long:"release" description:"Enable release mode (aka production) for app server"`
}

func Start(args []string) error {
	var opts Opts
	_, err := flags.ParseArgs(&opts, args)
	if err != nil {
		return err
	}

	request.StableVersions = opts.AvailableVersions.List
	fmt.Println("Running with V6 Stable Versions:", request.StableVersions)

	if opts.Release {
		gin.SetMode(gin.ReleaseMode)
	} else {
		fmt.Fprintf(os.Stderr, "warning: running in debug mode\n")
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()
	r.GET("/ping", request.Ping)
	r.GET("/edge", request.Edge)
	r.GET("/stable", request.Stable)
	r.Run(fmt.Sprintf(":%d", opts.Port))
	return nil
}
