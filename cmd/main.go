package main

import (
	"fmt"
	"os"

	"github.com/sagikazarmark/versio/pkg/cmd"
)

const (
	Version = "dev"
)

func main() {
	command := cmd.NewVersioCommand(Version)

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
