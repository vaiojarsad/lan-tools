package cmd

import (
	"github.com/spf13/cobra"

	"github.com/vaiojarsad/lan-tools/internal/services/dns"
)

func NewDNSSyncRecordsAForISPCommand() *cobra.Command {
	cmd := newDNSSyncRecordsAForISPCommand()
	return cmd
}

type dnsSyncRecordsAForISPParams struct {
	ispCode string
}

func newDNSSyncRecordsAForISPCommand() *cobra.Command {
	p := dnsSyncRecordsAForISPParams{}
	cmd := &cobra.Command{
		Use:   "sync-records-a-for-isp",
		Short: "Refreshes a type A dns records for a given isp",
		Long: "Refreshes a type A dns records for a given isp. A lookup for dns states associated to the input isp is made. " +
			"The refresh operation is carried out against the dns providers associated to the domains returned by the lookup. " +
			"All related local structures will be updated accordingly whenever necessary.",
		RunE: dnsSyncRecordsAForISPCommand(&p),
	}

	f := cmd.Flags()
	_ = cmd.MarkFlagRequired(addDNSStateISPCodeFlag(&p.ispCode, f))

	return cmd
}

func dnsSyncRecordsAForISPCommand(p *dnsSyncRecordsAForISPParams) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return dns.SyncRecordsA(p.ispCode)
	}
}
