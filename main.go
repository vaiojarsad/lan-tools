package main

import (
	"fmt"
	"os"

	"github.com/vaiojarsad/lan-tools/internal/cmd"
)

func main() {
	if err := cmd.NewLanToolsRootCommand().Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
