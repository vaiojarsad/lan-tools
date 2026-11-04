package cmd

import "github.com/spf13/cobra"

func NewCreateRootCommand() *cobra.Command {
	cmd := newCreateRootCommand()
	cmd.AddCommand(NewCreateDatabaseCommand())
	cmd.AddCommand(NewCreateDomainCommand())
	cmd.AddCommand(NewCreateISPCommand())
	cmd.AddCommand(NewCreateDNSProviderCommand())
	cmd.AddCommand(NewCreateDNSRecordACommand())
	return cmd
}

func newCreateRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates artifacts",
		Long:  "Creates artifacts in the local database and/or calling external services as necessary",
	}
	return cmd
}
