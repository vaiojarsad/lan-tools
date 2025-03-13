package cmd

import (
	"github.com/spf13/cobra"

	"github.com/vaiojarsad/lan-tools/internal/services/dns/state"
)

func NewCreateDnsStateCommand() *cobra.Command {
	cmd := newCreateDnsStateCommand()
	return cmd
}

type createDnsStateParams struct {
	domainName string
	ispCode    string
}

func newCreateDnsStateCommand() *cobra.Command {
	p := createDnsStateParams{}
	cmd := &cobra.Command{
		Use:   "dns-state",
		Short: "Creates a new \"state\" for a domain-isp tuple",
		Long: "Creates a new \"state\" for a domain-isp tuple. An entry for a domain/isp means that such domain is " +
			"accessible via the specified isp. The related last (locally) known dns provider state is stored here (" +
			"last public ip communicated to the provider, provider-specific dns' record id, sync state)",
		RunE: createDnsStateCommand(&p),
	}

	f := cmd.Flags()
	_ = cmd.MarkFlagRequired(addDnsStateDomainNameFlag(&p.domainName, f))
	_ = cmd.MarkFlagRequired(addDnsStateIspCodeFlag(&p.ispCode, f))

	return cmd
}

func createDnsStateCommand(p *createDnsStateParams) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return state.Create(p.domainName, p.ispCode)
	}
}
