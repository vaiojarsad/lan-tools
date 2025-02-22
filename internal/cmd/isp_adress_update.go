package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vaiojarsad/cloudflare-tools/internal/isp"
)

func NewIspAddressUpdateCommand() *cobra.Command {
	cmd := newIspAddressUpdateCommand()

	return cmd
}

func newIspAddressUpdateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update the locally stored by ISPs",
		Long:  "Update the locally stored by ISPs. By default tries to update the address for all ISPs",
		RunE:  ispAddressUpdateRun,
	}
	return cmd
}

func ispAddressUpdateRun(_ *cobra.Command, _ []string) error {
	return isp.UpdatePublicIP(nil)
}
