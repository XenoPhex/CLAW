package main

import (
	"fmt"
	"os"

	"code.cloudfoundry.org/claw/request"
	"github.com/gin-gonic/gin"
	flags "github.com/jessevdk/go-flags"
)

type Opts struct {
	Port    int  `long:"port" default:"8080" env:"PORT" description:"App server port"`
	Release bool `long:"release" description:"Enable release mode (aka production) for app server"`
}

func main() {
	var opts Opts
	_, err := flags.ParseArgs(&opts, os.Args)

	if err != nil {
		if _, ok := err.(*flags.Error); !ok {
			// GoFlags outputs it's own errors, don't need to double print.
			fmt.Fprintf(os.Stderr, err.Error())
		}
		os.Exit(1)
	}

	if opts.Release {
		gin.SetMode(gin.ReleaseMode)
	} else {
		fmt.Fprintf(os.Stderr, "warning: running in debug mode\n")
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()
	r.GET("/ping", request.Ping)
	r.Run(fmt.Sprintf(":%d", opts.Port))
}
