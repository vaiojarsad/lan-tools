package cmd

import (
	"github.com/spf13/cobra"
)

func NewISPRootCommand() *cobra.Command {
	cmd := newISPRootCommand()
	cmd.AddCommand(NewISPUpdatePublicIpCommand())
	return cmd
}

func newISPRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "isp",
		Short: "ISP management",
		Long:  "ISP management",
	}

	return cmd
}
