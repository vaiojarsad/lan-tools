package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewDnsListLocalCommand() *cobra.Command {
	cmd := newDnsListLocalCommand()
	return cmd
}

func newDnsListLocalCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "local",
		Short: "List DNS info in the local database",
		Long:  "List DNS info in the local database",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("dns list local")
			return nil
		},
	}
	return cmd
}
