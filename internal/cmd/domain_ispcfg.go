package cmd

import (
	"github.com/spf13/cobra"
)

func NewDomainIspCfgCommand() *cobra.Command {
	cmd := newDomainIspCfgCommand()
	cmd.AddCommand(NewDomainIspCfgCreateCommand())
	return cmd
}

func newDomainIspCfgCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ispcfg",
		Short: "Handles the isp configuration for domains",
		Long: "Manage the isp level configuration for a domain. An entry for a domain/isp means that such domain is " +
			"accessible via the specified isp. This also holds the information for the related dns provider (last " +
			"public ip set on the DNS and provider-specific dns' record id)",
	}

	return cmd
}
