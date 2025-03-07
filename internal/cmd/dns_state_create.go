package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vaiojarsad/lan-tools/internal/services/dns/state"
)

func NewDnsStateCreateCommand() *cobra.Command {
	cmd := newDnsStateCreateCommand()
	return cmd
}

type dnsStateCreateParams struct {
	domainName string
	ispCode    string
}

func newDnsStateCreateCommand() *cobra.Command {
	p := dnsStateCreateParams{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a new \"state\" for a domain-isp tuple",
		Long:  "Creates a new \"state\" for a domain-isp tuple",
		RunE:  domainIspCfgCreateCommand(&p),
	}

	f := cmd.Flags()
	_ = cmd.MarkFlagRequired(addDnsStateDomainNameFlag(&p.domainName, f))
	_ = cmd.MarkFlagRequired(addDnsStateIspCodeFlag(&p.ispCode, f))

	return cmd
}

func domainIspCfgCreateCommand(p *dnsStateCreateParams) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return state.Create(p.domainName, p.ispCode)
	}
}
