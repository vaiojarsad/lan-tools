package cmd

import (
	"github.com/spf13/cobra"

	"github.com/vaiojarsad/lan-tools/internal/database"
)

func NewCreateDatabaseCommand() *cobra.Command {
	cmd := newCreateDatabaseCommand()
	return cmd
}

func newCreateDatabaseCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "database",
		Short: "Creates the local database and its structures (tables)",
		Long:  "Creates the local database and its structures (tables)",
		RunE:  createDatabaseRun,
	}
	return cmd
}

func createDatabaseRun(_ *cobra.Command, _ []string) error {
	return database.Create()
}
