package config

type ISPConfig struct {
	Name               string
	PublicIPGetterType string
	PublicIPGetterCfg  map[string]string
	HostedDomains      []string
}
