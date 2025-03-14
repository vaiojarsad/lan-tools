package cmd

import (
	"github.com/spf13/cobra"

	"github.com/vaiojarsad/lan-tools/internal/services/dns"
)

func NewDnsSyncRecordsAForIspCommand() *cobra.Command {
	cmd := newDnsSyncRecordsAForIspCommand()
	return cmd
}

type dnsSyncRecordsAForIspParams struct {
	ispCode string
}

func newDnsSyncRecordsAForIspCommand() *cobra.Command {
	p := dnsSyncRecordsAForIspParams{}
	cmd := &cobra.Command{
		Use:   "sync-records-a-for-isp",
		Short: "Refreshes a type A dns records for a given isp",
		Long: "Refreshes a type A dns records for a given isp. A lookup for dns states associated to the input isp is made. " +
			"The refresh operation is carried out against the dns providers associated to the domains returned by the lookup. " +
			"All related local structures will be updated accordingly whenever necessary.",
		RunE: dnsSyncRecordsAForIspCommand(&p),
	}

	f := cmd.Flags()
	_ = cmd.MarkFlagRequired(addDnsStateIspCodeFlag(&p.ispCode, f))

	return cmd
}

func dnsSyncRecordsAForIspCommand(p *dnsSyncRecordsAForIspParams) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return dns.SyncRecordsA(p.ispCode)
	}
}
