package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/vaiojarsad/cloudflare-tools/internal/isp"
)

func NewISPRootCommand() *cobra.Command {
	cmd := newIspRootCommand()
	cmd.AddCommand(NewIspAddressCommand())

	return cmd
}

func newIspRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "isp",
		Short: "ISP interaction",
		Long: "ISP interaction. By default list ISPs and their current configuration and current locally saved public " +
			"IP",
		RunE: ispRootRun,
	}
	return cmd
}

func ispRootRun(_ *cobra.Command, _ []string) error {
	ispList, err := isp.List()
	if err != nil {
		return err
	}
	for _, ispData := range ispList {
		fmt.Printf("%s\n", ispData)
	}
	return nil
}
