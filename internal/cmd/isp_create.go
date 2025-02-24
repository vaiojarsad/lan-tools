package cmd

import (
	//"fmt"

	"github.com/spf13/cobra"
	"github.com/vaiojarsad/lan-tools/internal/isp"
	//"github.com/vaiojarsad/lan-tools/internal/environment"
	//"github.com/vaiojarsad/lan-tools/internal/isp"
)

func NewIspCreateCommand() *cobra.Command {
	cmd := newIspCreateCommand()
	return cmd
}

type ispCreateParams struct {
	code               string
	name               string
	publicIpGetterType string
	publicIpGetterCfg  map[string]string
}

func newIspCreateCommand() *cobra.Command {
	p := ispCreateParams{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Inserts the new ISP",
		Long:  "Inserts the new ISP",
		RunE:  ispCreateCommand(&p),
	}
	f := cmd.Flags()
	_ = cmd.MarkFlagRequired(addISPCodeFlag(&p.code, f))
	_ = cmd.MarkFlagRequired(addISPNameFlag(&p.name, f))
	addPublicIpGetterTypeFlag(&p.publicIpGetterType, f)
	addPublicIpGetterCfgFlag(&p.publicIpGetterCfg, f)
	return cmd
}

func ispCreateCommand(p *ispCreateParams) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return isp.Create(p.code, p.name, p.publicIpGetterType, p.publicIpGetterCfg)
	}
}
