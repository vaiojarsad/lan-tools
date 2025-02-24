package cmd

import "github.com/spf13/pflag"

// Globals & commons
func addPublicIpGetterTypeFlag(s *string, f *pflag.FlagSet) {
	f.StringVar(s, publicIpGetterType, "", "specifies a public ip getter type (ex. ipify)")
}

func addPublicIpGetterCfgFlag(s *map[string]string, f *pflag.FlagSet) {
	f.StringToStringVar(s, publicIpGetterCfg, map[string]string{}, "specifies the public ip getter configuration "+
		"as a string to string map")
}

// ISP
func addISPCodeFlag(s *string, f *pflag.FlagSet) string {
	f.StringVarP(s, ispCodeFlag, ispCodeShortHandFlag, "", "specifies unique acronym that identifies the ISP")
	return ispCodeFlag
}

func addISPNameFlag(s *string, f *pflag.FlagSet) string {
	f.StringVarP(s, ispNameFlag, ispNameShortHandFlag, "", "specifies the ISP name")
	return ispNameFlag
}
