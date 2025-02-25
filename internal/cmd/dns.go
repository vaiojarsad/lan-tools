package cmd

import "github.com/spf13/cobra"

func NewDnsRootCommand() *cobra.Command {
	cmd := newDnsRootCommand()
	cmd.AddCommand(NewDnsProviderRootCommand())

	return cmd
}

func newDnsRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dns",
		Short: "Cloudflare DNS api interaction",
		Long:  "Cloudflare DNS api interaction",
	}
	return cmd
}
