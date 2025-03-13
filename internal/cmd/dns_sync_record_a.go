package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vaiojarsad/lan-tools/internal/services/dns"
)

func NewDnsSyncRecordACommand() *cobra.Command {
	cmd := newDnsSyncRecordACommand()
	return cmd
}

type dnsSyncRecordAParams struct {
	domainName string
	ispCode    string
}

func newDnsSyncRecordACommand() *cobra.Command {
	p := dnsSyncRecordAParams{}
	cmd := &cobra.Command{
		Use:   "sync-record-a",
		Short: "Refreshes a type A dns record for a domain-isp tuple",
		Long: "Refreshes a type A dns record for a domain-isp tuple. The operation is carried out against the dns provider " +
			"associated to the domain. All related local structures will be updated accordingly whenever necessary. " +
			"Local dns state for the domain-isp must exists.",
		RunE: dnsSyncRecordACommand(&p),
	}

	f := cmd.Flags()
	_ = cmd.MarkFlagRequired(addDnsStateDomainNameFlag(&p.domainName, f))
	_ = cmd.MarkFlagRequired(addDnsStateIspCodeFlag(&p.ispCode, f))

	return cmd
}

func dnsSyncRecordACommand(p *dnsSyncRecordAParams) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return dns.SyncRecordA(p.domainName, p.ispCode)
	}
}
