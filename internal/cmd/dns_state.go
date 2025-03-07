package cmd

import (
	"github.com/spf13/cobra"
)

func NewDnsStateRootCommand() *cobra.Command {
	cmd := newDnsStateRootCommand()
	cmd.AddCommand(NewDnsStateCreateCommand())
	return cmd
}

func newDnsStateRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "state",
		Short: "Handles the dns state for a domain-isp tuple",
		Long: "Manage the dns state for a domain-isp. An entry for a domain/isp means that such domain is " +
			"accessible via the specified isp. The related last (locally) known dns provider state is stored here (" +
			"last public ip communicated to the provider, provider-specific dns' record id, sync state)",
	}

	return cmd
}
