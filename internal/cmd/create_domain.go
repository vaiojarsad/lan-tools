package cmd

import (
	"github.com/spf13/cobra"

	"github.com/vaiojarsad/lan-tools/internal/services/domain"
)

func NewCreateDomainCommand() *cobra.Command {
	cmd := newCreateDomainCommand()
	return cmd
}

type createDomainParams struct {
	name            string
	description     string
	dnsProviderCode string
}

func newCreateDomainCommand() *cobra.Command {
	p := createDomainParams{}
	cmd := &cobra.Command{
		Use:   "domain",
		Short: "Creates a new domain",
		Long:  "Creates a new domain",
		RunE:  createDomainCommand(&p),
	}

	f := cmd.Flags()
	_ = cmd.MarkFlagRequired(addDomainNameFlag(&p.name, f))
	_ = cmd.MarkFlagRequired(addDomainDnsProviderCodeFlag(&p.dnsProviderCode, f))
	addDomainDescriptionFlag(&p.description, f)

	return cmd
}

func createDomainCommand(p *createDomainParams) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return domain.Create(p.name, p.description, p.dnsProviderCode)
	}
}
