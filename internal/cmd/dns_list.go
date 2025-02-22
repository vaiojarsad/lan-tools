package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewDnsListRootCommand() *cobra.Command {
	cmd := newDnsListRootCommand()
	cmd.AddCommand(NewDnsListLocalCommand())
	cmd.AddCommand(NewDnsListRemoteCommand())
	return cmd
}

func newDnsListRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List DNS from both, the local database and CloudFlare",
		Long:  "List DNS from both, the local database and CloudFlare",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("dns list")
			return nil
		},
	}

	return cmd
}
