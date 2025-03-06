package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vaiojarsad/lan-tools/internal/services/domain"
)

func NewDomainCreateCommand() *cobra.Command {
	cmd := newDomainCreateCommand()
	return cmd
}

type domainCreateParams struct {
	name            string
	description     string
	dnsProviderCode string
}

func newDomainCreateCommand() *cobra.Command {
	p := domainCreateParams{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a new domain",
		Long:  "Creates a new domain",
		RunE:  domainCreateCommand(&p),
	}

	f := cmd.Flags()
	_ = cmd.MarkFlagRequired(addDomainNameFlag(&p.name, f))
	_ = cmd.MarkFlagRequired(addDomainDNSProviderCodeFlag(&p.dnsProviderCode, f))
	addDomainDescriptionFlag(&p.description, f)

	return cmd
}

func domainCreateCommand(p *domainCreateParams) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return domain.Create(p.name, p.description, p.dnsProviderCode)
	}
}
