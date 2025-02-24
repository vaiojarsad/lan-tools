package cmd

import (
	//"fmt"

	"github.com/spf13/cobra"
	"github.com/vaiojarsad/lan-tools/internal/isp"
	//"github.com/vaiojarsad/lan-tools/internal/environment"
	//"github.com/vaiojarsad/lan-tools/internal/isp"
)

func NewIspRefreshPublicIpCommand() *cobra.Command {
	cmd := newIspRefreshPublicIpCommand()
	return cmd
}

type ispRefreshPublicIpParams struct {
	code string
}

func newIspRefreshPublicIpCommand() *cobra.Command {
	p := ispRefreshPublicIpParams{}
	cmd := &cobra.Command{
		Use:   "refresh-public-ip",
		Short: "Refresh ISP's public IP",
		Long:  "Gets the public ip from the ISP and saves it to the local DB",
		RunE:  ispRefreshPublicIpCommand(&p),
	}
	f := cmd.Flags()
	_ = cmd.MarkFlagRequired(addISPCodeFlag(&p.code, f))
	return cmd
}

func ispRefreshPublicIpCommand(p *ispRefreshPublicIpParams) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return isp.RefreshIspPublicIp(p.code)
	}
}
