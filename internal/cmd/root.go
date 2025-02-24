// Package cmd is used to define Cobra stuff
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vaiojarsad/lan-tools/internal/utils/loggerutils"

	"github.com/vaiojarsad/lan-tools/internal/config"
	"github.com/vaiojarsad/lan-tools/internal/environment"
	"github.com/vaiojarsad/lan-tools/internal/utils"
)

func NewCloudFlareToolsRootCommand() *cobra.Command {
	cobra.MousetrapHelpText = ""
	cmd := newCloudFlareToolsRootCommand()
	cmd.AddCommand(NewDnsRootCommand())
	cmd.AddCommand(NewISPRootCommand())
	cmd.AddCommand(NewDatabaseRootCommand())

	utils.ForEach(cmd.Commands(), func(c *cobra.Command) {
		c.PersistentFlags().String(cfgFileFlag, "", "specifies the configuration file")
		c.PersistentPreRunE = loadCfg
	})

	return cmd
}

func newCloudFlareToolsRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lan-tools",
		Short: "Cloudflare toolset",
		Long:  "Cloudflare toolset that interacts with the Cloudflare API",
	}
	return cmd
}

func loadCfg(c *cobra.Command, _ []string) error {
	configFile, err := c.Flags().GetString(cfgFileFlag)
	if err != nil {
		configFile = ""
	}
	env := environment.Create()
	env.ErrorLogger = loggerutils.GetStdErrorLogger()
	env.OutputLogger = loggerutils.GetStdOutputLogger()
	cm, err := config.Create(configFile)
	if err != nil {
		return err
	}
	env.ConfigManager = cm
	return nil
}
