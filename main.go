package main

import (
	"fmt"
	"os"

	cmd "github.com/irevenko/weebsay/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
