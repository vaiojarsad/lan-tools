package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewDnsProviderRootCommand() *cobra.Command {
	cmd := newDnsProviderRootCommand()
	cmd.AddCommand(NewDnsProviderCreateCommand())
	return cmd
}

func newDnsProviderRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "provider",
		Short: "Manages DNS providers",
		Long:  "Manages DNS providers",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("dns list")
			return nil
		},
	}

	return cmd
}
