// Package cmd is used to define Cobra stuff
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/vaiojarsad/lan-tools/internal/config"
	"github.com/vaiojarsad/lan-tools/internal/environment"
	"github.com/vaiojarsad/lan-tools/internal/events"
	dnsevents "github.com/vaiojarsad/lan-tools/internal/services/dns/events"
	"github.com/vaiojarsad/lan-tools/internal/utils"
	"github.com/vaiojarsad/lan-tools/internal/utils/loggerutils"
)

func NewLanToolsRootCommand() *cobra.Command {
	cobra.MousetrapHelpText = ""
	cmd := newLanToolsRootCommand()
	cmd.AddCommand(NewDNSRootCommand())
	cmd.AddCommand(NewISPRootCommand())
	cmd.AddCommand(NewCreateRootCommand())

	utils.ForEach(cmd.Commands(), func(c *cobra.Command) {
		c.PersistentFlags().String(cfgFileFlag, "", "specifies the configuration file")
		c.PersistentPreRunE = loadCfg
	})

	return cmd
}

func newLanToolsRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lan-tools",
		Short: "lan toolset",
		Long:  "lan toolset",
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
	env.IPChangedPublisher = events.NewInMemoryBus[*events.IPChanged]()
	env.IPChangedPublisher.Subscribe(dnsevents.HandleISPPublicIPChanged)
	cm, err := config.Create(configFile)
	if err != nil {
		return err
	}
	env.ConfigManager = cm
	return nil
}
