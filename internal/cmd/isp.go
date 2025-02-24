package cmd

import (
	"github.com/spf13/cobra"
)

func NewISPRootCommand() *cobra.Command {
	cmd := newIspRootCommand()
	cmd.AddCommand(NewIspCreateCommand())
	cmd.AddCommand(NewIspRefreshPublicIpCommand())
	return cmd
}

func newIspRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "isp",
		Short: "ISP interaction",
		Long:  "ISP interaction",
	}

	return cmd
}
