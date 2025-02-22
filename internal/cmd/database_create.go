package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vaiojarsad/cloudflare-tools/internal/database"
)

func NewDatabaseCreateCommand() *cobra.Command {
	cmd := newDatabaseCreateCommand()
	return cmd
}

func newDatabaseCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates the local database and its structures (tables)",
		Long:  "Creates the local database and its structures (tables)",
		RunE:  databaseCreateRun,
	}
	return cmd
}

func databaseCreateRun(_ *cobra.Command, _ []string) error {
	return database.Create()
}
