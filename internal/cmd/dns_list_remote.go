package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewDnsListRemoteCommand() *cobra.Command {
	cmd := newDnsListRemoteCommand()
	return cmd
}

func newDnsListRemoteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remote",
		Short: "List DNS info from Cloudflare",
		Long:  "List DNS info from CloudFlare",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("dns list remote")
			return nil
		},
	}
	return cmd
}
