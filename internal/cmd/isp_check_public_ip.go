package cmd

import (
	//"fmt"

	"github.com/spf13/cobra"
	"github.com/vaiojarsad/lan-tools/internal/services/isp"
	//"github.com/vaiojarsad/lan-tools/internal/environment"
	//"github.com/vaiojarsad/lan-tools/internal/isp"
)

func NewISPCheckPublicIpCommand() *cobra.Command {
	cmd := newISPCheckPublicIpCommand()
	return cmd
}

type ispCheckPublicIpParams struct {
	code    string
	trigger bool
}

func newISPCheckPublicIpCommand() *cobra.Command {
	p := ispCheckPublicIpParams{}
	cmd := &cobra.Command{
		Use:   "check-public-ip",
		Short: "Check ISP's public IP",
		Long:  "Check the public ip from the ISP",
		RunE:  ispCheckPublicIpCommand(&p),
	}
	f := cmd.Flags()
	_ = cmd.MarkFlagRequired(addISPCodeFlag(&p.code, f))
	_ = cmd.MarkFlagRequired(addTriggerRelatedFlag(&p.trigger, f))
	return cmd
}

func ispCheckPublicIpCommand(p *ispCheckPublicIpParams) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return isp.CheckISPPublicIp(p.code, p.trigger)
	}
}
