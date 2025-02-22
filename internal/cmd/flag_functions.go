package cmd

import "github.com/spf13/pflag"

func addCfgFileFlag(s *string, f *pflag.FlagSet) {
	f.StringVarP(s, cfgFileFlag, cfgFileShortHandFlag, "", "specifies the configuration file")
}
