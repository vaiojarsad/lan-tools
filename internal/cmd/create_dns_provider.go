package cmd

import (
	"github.com/spf13/cobra"

	"github.com/vaiojarsad/lan-tools/internal/services/dns/provider"
)

func NewCreateDNSProviderCommand() *cobra.Command {
	cmd := newCreateDNSProviderCommand()
	return cmd
}

type createDNSProviderParams struct {
	code        string
	name        string
	serviceType string
	serviceCfg  map[string]string
}

func newCreateDNSProviderCommand() *cobra.Command {
	p := createDNSProviderParams{}
	cmd := &cobra.Command{
		Use:   "dns-provider",
		Short: "Creates a new DNS provider",
		Long:  "Creates a new DNS provider",
		RunE:  createDNSProviderCommand(&p),
	}
	f := cmd.Flags()
	_ = cmd.MarkFlagRequired(addDNSProviderCodeFlag(&p.code, f))
	_ = cmd.MarkFlagRequired(addDNSProviderNameFlag(&p.name, f))
	addDNSProviderServiceTypeFlag(&p.serviceType, f)
	addDNSProviderServiceCfgFlag(&p.serviceCfg, f)
	return cmd
}

func createDNSProviderCommand(p *createDNSProviderParams) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return provider.Create(p.code, p.name, p.serviceType, p.serviceCfg)
	}
}
