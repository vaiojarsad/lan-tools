package cmd

import "github.com/spf13/pflag"

// Globals & commons

// ISP
func addIspCodeFlag(s *string, f *pflag.FlagSet) string {
	f.StringVarP(s, ispCodeFlag, providerCodeShortHandFlag, "", "specifies unique acronym that identifies the ISP")
	return ispCodeFlag
}

func addIspNameFlag(s *string, f *pflag.FlagSet) string {
	f.StringVarP(s, nameFlag, nameShortHandFlag, "", "specifies the ISP name")
	return nameFlag
}

func addPublicIpGetterTypeFlag(s *string, f *pflag.FlagSet) {
	f.StringVar(s, publicIpGetterType, "", "specifies a public ip getter type (ex. ipify)")
}

func addPublicIpGetterCfgFlag(s *map[string]string, f *pflag.FlagSet) {
	f.StringToStringVar(s, publicIpGetterCfg, map[string]string{}, "specifies the public ip getter configuration "+
		"as a string to string map")
}

// DNS Provider
func addDnsProviderCodeFlag(s *string, f *pflag.FlagSet) string {
	f.StringVarP(s, dnsProviderCodeFlag, providerCodeShortHandFlag, "", "specifies unique acronym that identifies the DNS provider")
	return dnsProviderCodeFlag
}

func addDnsProviderNameFlag(s *string, f *pflag.FlagSet) string {
	f.StringVarP(s, nameFlag, nameShortHandFlag, "", "specifies the DNS provider name")
	return nameFlag
}

func addDnsProviderServiceTypeFlag(s *string, f *pflag.FlagSet) {
	f.StringVar(s, dnsProviderServiceType, "", "specifies a DNS service provider type (ex. Cloudflare)")
}

func addDnsProviderServiceCfgFlag(s *map[string]string, f *pflag.FlagSet) {
	f.StringToStringVar(s, dnsProviderServiceCfg, map[string]string{}, "specifies the DNS service provider configuration "+
		"as a string to string map")
}

// Domain
func addDomainNameFlag(s *string, f *pflag.FlagSet) string {
	f.StringVarP(s, nameFlag, nameShortHandFlag, "", "specifies the domain name")
	return nameFlag
}

func addDomainDescriptionFlag(s *string, f *pflag.FlagSet) string {
	f.StringVarP(s, descriptionFlag, descriptionShortHandFlag, "", "specifies a brief description about the domain")
	return descriptionFlag
}

func addDomainDnsProviderCodeFlag(s *string, f *pflag.FlagSet) string {
	f.StringVarP(s, dnsProviderCodeFlag, providerCodeShortHandFlag, "", "specifies the code for the DNS provider that resolves this domain")
	return dnsProviderCodeFlag
}

// DNS State
func addDnsStateDomainNameFlag(s *string, f *pflag.FlagSet) string {
	f.StringVarP(s, domainNameFlag, domainNameShortHandFlag, "", "specifies the domain name")
	return domainNameFlag
}

func addDnsStateIspCodeFlag(s *string, f *pflag.FlagSet) string {
	f.StringVarP(s, ispCodeFlag, codeShortHandFlag, "", "specifies the code for the ISP we want to associate the domain with")
	return ispCodeFlag
}
