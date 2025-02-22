package cmd

import "github.com/spf13/cobra"

func NewDatabaseRootCommand() *cobra.Command {
	cmd := newDatabaseRootCommand()
	cmd.AddCommand(NewDatabaseCreateCommand())

	return cmd
}

func newDatabaseRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "database",
		Short: "Local database interaction",
		Long:  "Local database interaction",
	}
	return cmd
}
