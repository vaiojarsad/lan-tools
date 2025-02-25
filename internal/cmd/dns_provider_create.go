package cmd

import (
	"github.com/spf13/cobra"

	"github.com/vaiojarsad/lan-tools/internal/dns/provider"
)

func NewDnsProviderCreateCommand() *cobra.Command {
	cmd := newDnsProviderCreateCommand()
	return cmd
}

type dnsProviderCreateParams struct {
	code        string
	name        string
	serviceType string
	serviceCfg  map[string]string
}

func newDnsProviderCreateCommand() *cobra.Command {
	p := dnsProviderCreateParams{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a new DNS provider",
		Long:  "Creates a new DNS provider",
		RunE:  dnsProviderCreateCommand(&p),
	}
	f := cmd.Flags()
	_ = cmd.MarkFlagRequired(addDNSProviderCodeFlag(&p.code, f))
	_ = cmd.MarkFlagRequired(addDNSProviderNameFlag(&p.name, f))
	addDNSProviderServiceTypeFlag(&p.serviceType, f)
	addDNSProviderServiceCfgFlag(&p.serviceCfg, f)
	return cmd
}

func dnsProviderCreateCommand(p *dnsProviderCreateParams) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return provider.Create(p.code, p.name, p.serviceType, p.serviceCfg)
	}
}
