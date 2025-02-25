package cmd

import "github.com/spf13/pflag"

// Globals & commons

// ISP
func addISPCodeFlag(s *string, f *pflag.FlagSet) string {
	f.StringVarP(s, ispCodeFlag, ispCodeShortHandFlag, "", "specifies unique acronym that identifies the ISP")
	return ispCodeFlag
}

func addISPNameFlag(s *string, f *pflag.FlagSet) string {
	f.StringVarP(s, ispNameFlag, ispNameShortHandFlag, "", "specifies the ISP name")
	return ispNameFlag
}

func addPublicIpGetterTypeFlag(s *string, f *pflag.FlagSet) {
	f.StringVar(s, publicIpGetterType, "", "specifies a public ip getter type (ex. ipify)")
}

func addPublicIpGetterCfgFlag(s *map[string]string, f *pflag.FlagSet) {
	f.StringToStringVar(s, publicIpGetterCfg, map[string]string{}, "specifies the public ip getter configuration "+
		"as a string to string map")
}

// DNS Provider
func addDNSProviderCodeFlag(s *string, f *pflag.FlagSet) string {
	f.StringVarP(s, dnsProviderCodeFlag, dnsProviderCodeShortHandFlag, "", "specifies unique acronym that identifies the DNS provider")
	return dnsProviderCodeFlag
}

func addDNSProviderNameFlag(s *string, f *pflag.FlagSet) string {
	f.StringVarP(s, dnsProviderNameFlag, dnsProviderNameShortHandFlag, "", "specifies the DNS provider name")
	return dnsProviderNameFlag
}

func addDNSProviderServiceTypeFlag(s *string, f *pflag.FlagSet) {
	f.StringVar(s, dnsProviderServiceType, "", "specifies a DNS service provider type (ex. Cloudflare)")
}

func addDNSProviderServiceCfgFlag(s *map[string]string, f *pflag.FlagSet) {
	f.StringToStringVar(s, dnsProviderServiceCfg, map[string]string{}, "specifies the DNS service provider configuration "+
		"as a string to string map")
}
