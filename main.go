package main

import (
	"fmt"
	"os"

	"code.cloudfoundry.org/claw/exec"
)

func main() {
	server, listenAddr, err := exec.Server(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	server.Run(listenAddr)
}
