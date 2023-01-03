package main

import (
	"os"
	"github.com/Ali15700/assignment02bca/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()

}
