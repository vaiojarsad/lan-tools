package cmd

import (
	"github.com/spf13/cobra"

	"github.com/vaiojarsad/lan-tools/internal/services/dns"
)

func NewCreateDNSRecordACommand() *cobra.Command {
	cmd := newCreateDNSRecordACommand()
	return cmd
}

type createDNSRecordAParams struct {
	domainName string
	ispCode    string
}

func newCreateDNSRecordACommand() *cobra.Command {
	p := createDNSRecordAParams{}
	cmd := &cobra.Command{
		Use:   "dns-record-a",
		Short: "Creates a type A dns record for a domain-isp tuple",
		Long: "Creates a type A dns record for a domain-isp tuple. The operation is carried out against the dns provider " +
			"associated to the domain. All related local structures will be updated accordingly whenever necessary. " +
			"Even if there is no local dns state for the domain-isp, we assume that might already exists a related record " +
			"on provider side.",
		RunE: createDNSRecordACommand(&p),
	}

	f := cmd.Flags()
	_ = cmd.MarkFlagRequired(addDNSStateDomainNameFlag(&p.domainName, f))
	_ = cmd.MarkFlagRequired(addDNSStateISPCodeFlag(&p.ispCode, f))

	return cmd
}

func createDNSRecordACommand(p *createDNSRecordAParams) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return dns.CreateRecordA(p.domainName, p.ispCode)
	}
}
