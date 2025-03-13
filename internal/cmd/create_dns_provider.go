package cmd

import (
	"github.com/spf13/cobra"

	"github.com/vaiojarsad/lan-tools/internal/services/dns/provider"
)

func NewCreateDnsProviderCommand() *cobra.Command {
	cmd := newCreateDnsProviderCommand()
	return cmd
}

type createDnsProviderParams struct {
	code        string
	name        string
	serviceType string
	serviceCfg  map[string]string
}

func newCreateDnsProviderCommand() *cobra.Command {
	p := createDnsProviderParams{}
	cmd := &cobra.Command{
		Use:   "dns-provider",
		Short: "Creates a new DNS provider",
		Long:  "Creates a new DNS provider",
		RunE:  createDnsProviderCommand(&p),
	}
	f := cmd.Flags()
	_ = cmd.MarkFlagRequired(addDnsProviderCodeFlag(&p.code, f))
	_ = cmd.MarkFlagRequired(addDnsProviderNameFlag(&p.name, f))
	addDnsProviderServiceTypeFlag(&p.serviceType, f)
	addDnsProviderServiceCfgFlag(&p.serviceCfg, f)
	return cmd
}

func createDnsProviderCommand(p *createDnsProviderParams) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return provider.Create(p.code, p.name, p.serviceType, p.serviceCfg)
	}
}
