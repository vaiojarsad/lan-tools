package cmd

import "github.com/spf13/cobra"

func NewDomainRootCommand() *cobra.Command {
	cmd := newDomainRootCommand()
	cmd.AddCommand(NewDomainCreateCommand())
	cmd.AddCommand(NewDomainIspCfgCommand())

	return cmd
}

func newDomainRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "domain",
		Short: "Domain level utilities",
		Long:  "Domain level utilities",
	}
	return cmd
}
