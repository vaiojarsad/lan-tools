package cmd

import (
	//"fmt"

	"github.com/spf13/cobra"
	"github.com/vaiojarsad/lan-tools/internal/services/isp"
	//"github.com/vaiojarsad/lan-tools/internal/environment"
	//"github.com/vaiojarsad/lan-tools/internal/isp"
)

func NewISPUpdatePublicIpCommand() *cobra.Command {
	cmd := newISPUpdatePublicIpCommand()
	return cmd
}

type ispUpdatePublicIpParams struct {
	code    string
	trigger bool
}

func newISPUpdatePublicIpCommand() *cobra.Command {
	p := ispUpdatePublicIpParams{}
	cmd := &cobra.Command{
		Use:   "update-public-ip",
		Short: "Update ISP's public IP",
		Long:  "Gets the public ip from the ISP and saves it to the local DB",
		RunE:  ispUpdatePublicIpCommand(&p),
	}
	f := cmd.Flags()
	_ = cmd.MarkFlagRequired(addISPCodeFlag(&p.code, f))
	// _ = cmd.MarkFlagRequired(addTriggerRelatedFlag(&p.trigger, f))
	return cmd
}

func ispUpdatePublicIpCommand(p *ispUpdatePublicIpParams) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return isp.UpdateISPPublicIp(p.code)
	}
}
