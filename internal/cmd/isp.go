package cmd

import (
	"github.com/spf13/cobra"
)

func NewISPRootCommand() *cobra.Command {
	cmd := newIspRootCommand()
	cmd.AddCommand(NewIspRefreshPublicIpCommand())
	return cmd
}

func newIspRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "isp",
		Short: "ISP management",
		Long:  "ISP management",
	}

	return cmd
}
