package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vaiojarsad/lan-tools/internal/domain/ispcfg"
)

func NewDomainIspCfgCreateCommand() *cobra.Command {
	cmd := newDomainIspCfgCreateCommand()
	return cmd
}

type domainIspCfgCreateParams struct {
	domainName string
	ispCode    string
}

func newDomainIspCfgCreateCommand() *cobra.Command {
	p := domainIspCfgCreateParams{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a new isp level configuration for a domain",
		Long:  "Creates a new isp level configuration for a domain",
		RunE:  domainIspCfgCreateCommand(&p),
	}

	f := cmd.Flags()
	_ = cmd.MarkFlagRequired(addDomainIspCfgDomainNameFlag(&p.domainName, f))
	_ = cmd.MarkFlagRequired(addDomainIspCfgIspCodeFlag(&p.ispCode, f))

	return cmd
}

func domainIspCfgCreateCommand(p *domainIspCfgCreateParams) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return ispcfg.Create(p.domainName, p.ispCode)
	}
}
