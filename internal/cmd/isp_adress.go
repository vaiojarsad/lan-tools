package cmd

import (
	//"fmt"

	"github.com/spf13/cobra"
	//"github.com/vaiojarsad/cloudflare-tools/internal/environment"
	//"github.com/vaiojarsad/cloudflare-tools/internal/isp"
)

func NewIspAddressCommand() *cobra.Command {
	cmd := newIspAddressCommand()
	cmd.AddCommand(NewIspAddressUpdateCommand())

	return cmd
}

func newIspAddressCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "address",
		Short: "ISP public address interaction",
		Long:  "ISP public address interaction",
	}
	return cmd
}
