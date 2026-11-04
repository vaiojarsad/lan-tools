package cmd

import "github.com/spf13/cobra"

func NewDNSRootCommand() *cobra.Command {
	cmd := newDNSRootCommand()
	cmd.AddCommand(NewDNSSyncRecordACommand())
	cmd.AddCommand(NewDNSSyncRecordsAForISPCommand())
	return cmd
}

func newDNSRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dns",
		Short: "Domain name service utilities",
		Long:  "Domain name service utilities",
	}
	return cmd
}
