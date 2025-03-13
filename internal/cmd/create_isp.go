package cmd

import (
	"github.com/spf13/cobra"

	"github.com/vaiojarsad/lan-tools/internal/services/isp"
)

func NewCreateIspCommand() *cobra.Command {
	cmd := newCreateIspCommand()
	return cmd
}

type createIspParams struct {
	code               string
	name               string
	publicIpGetterType string
	publicIpGetterCfg  map[string]string
}

func newCreateIspCommand() *cobra.Command {
	p := createIspParams{}
	cmd := &cobra.Command{
		Use:   "isp",
		Short: "Creates a new ISP",
		Long:  "Creates a new ISP",
		RunE:  ispCreateCommand(&p),
	}
	f := cmd.Flags()
	_ = cmd.MarkFlagRequired(addIspCodeFlag(&p.code, f))
	_ = cmd.MarkFlagRequired(addIspNameFlag(&p.name, f))
	addPublicIpGetterTypeFlag(&p.publicIpGetterType, f)
	addPublicIpGetterCfgFlag(&p.publicIpGetterCfg, f)
	return cmd
}

func ispCreateCommand(p *createIspParams) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return isp.Create(p.code, p.name, p.publicIpGetterType, p.publicIpGetterCfg)
	}
}
