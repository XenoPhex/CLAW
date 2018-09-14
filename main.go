package main

import (
	"fmt"
	"os"

	"code.cloudfoundry.org/claw/exec"
	flags "github.com/jessevdk/go-flags"
)

func main() {
	err := exec.Start(os.Args)
	if err != nil {
		if _, ok := err.(*flags.Error); !ok {
			// GoFlags outputs it's own errors, don't need to double print.
			fmt.Fprintf(os.Stderr, err.Error())
		}
		os.Exit(1)
	}
}
