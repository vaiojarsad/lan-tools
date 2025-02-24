package main

import (
	"fmt"
	"github.com/vaiojarsad/lan-tools/internal/cmd"
	"os"
)

func main() {
	if err := cmd.NewCloudFlareToolsRootCommand().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
