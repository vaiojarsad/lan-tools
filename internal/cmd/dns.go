package cmd

import "github.com/spf13/cobra"

func NewDnsRootCommand() *cobra.Command {
	cmd := newDnsRootCommand()
	cmd.AddCommand(NewDnsProviderRootCommand())
	cmd.AddCommand(NewDnsStateRootCommand())

	return cmd
}

func newDnsRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dns",
		Short: "Domain name service utilities",
		Long:  "Domain name service utilities",
	}
	return cmd
}
