// Package cmd is used to define Cobra stuff
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vaiojarsad/cloudflare-tools/internal/utils/loggerutils"

	"github.com/vaiojarsad/cloudflare-tools/internal/config"
	"github.com/vaiojarsad/cloudflare-tools/internal/environment"
	"github.com/vaiojarsad/cloudflare-tools/internal/utils"
)

func NewCloudFlareToolsRootCommand() *cobra.Command {
	cobra.MousetrapHelpText = ""
	cmd := newCloudFlareToolsRootCommand()
	cmd.AddCommand(NewDnsRootCommand())
	cmd.AddCommand(NewISPRootCommand())
	cmd.AddCommand(NewDatabaseRootCommand())

	utils.ForEach(cmd.Commands(), func(c *cobra.Command) {
		c.PersistentPreRunE = loadCfg
	})

	return cmd
}

func newCloudFlareToolsRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cloudflare-tools",
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
	env.ConfigFile = configFile
	env.ErrorLogger = loggerutils.GetStdErrorLogger()
	env.OutputLogger = loggerutils.GetStdOutputLogger()
	cm, err := config.Create(configFile)
	if err != nil {
		return err
	}
	env.ConfigManager = cm
	return nil
}
