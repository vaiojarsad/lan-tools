package cmd

import (
	"github.com/spf13/cobra"

	"github.com/vaiojarsad/lan-tools/internal/services/isp"
)

func NewCreateISPCommand() *cobra.Command {
	cmd := newCreateISPCommand()
	return cmd
}

type createISPParams struct {
	code               string
	name               string
	publicIpGetterType string
	publicIpGetterCfg  map[string]string
}

func newCreateISPCommand() *cobra.Command {
	p := createISPParams{}
	cmd := &cobra.Command{
		Use:   "isp",
		Short: "Creates a new ISP",
		Long:  "Creates a new ISP",
		RunE:  ispCreateCommand(&p),
	}
	f := cmd.Flags()
	_ = cmd.MarkFlagRequired(addISPCodeFlag(&p.code, f))
	_ = cmd.MarkFlagRequired(addISPNameFlag(&p.name, f))
	addPublicIpGetterTypeFlag(&p.publicIpGetterType, f)
	addPublicIpGetterCfgFlag(&p.publicIpGetterCfg, f)
	return cmd
}

func ispCreateCommand(p *createISPParams) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return isp.Create(p.code, p.name, p.publicIpGetterType, p.publicIpGetterCfg)
	}
}
