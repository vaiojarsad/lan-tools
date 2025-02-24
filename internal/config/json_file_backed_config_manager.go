package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type jsonISPsConfig struct {
	Name               string            `json:"name,omitempty"`
	PublicIPGetterType string            `json:"public_ip_getter_type,omitempty"`
	PublicIPGetterCfg  map[string]string `json:"public_ip_getter_cfg,omitempty"`
	HostedDomains      []string          `json:"hosted_domains,omitempty"`
}

type jsonSMTPConfig struct {
	Host   string `json:"host,omitempty"`
	Port   int    `json:"port,omitempty"`
	Sender string `json:"sender,omitempty"`
	Pass   string `json:"password,omitempty"`
	To     string `json:"to,omitempty"`
}

type jsonDatabaseConfig struct {
	Path string `json:"path,omitempty"`
	Name string `json:"name,omitempty"`
}

type jsonCloudflareConfig struct {
	Token   string   `json:"token,omitempty"`
	Domains []string `json:"domains,omitempty"`
}

type jsonConfig struct {
	SMTPConfig       jsonSMTPConfig            `json:"smtp_config,omitempty"`
	ISPsConfig       map[string]jsonISPsConfig `json:"isps_config,omitempty"`
	DatabaseConfig   jsonDatabaseConfig        `json:"database_config,omitempty"`
	CloudflareConfig jsonCloudflareConfig      `json:"cloudflare_config,omitempty"`
}

type jsonFileBackedConfigManager struct {
	smtpConfig       *SMTPConfig
	databaseConfig   *DatabaseConfig
	cloudflareConfig *CloudflareConfig
	ispsConfig       map[string]*ISPConfig
}

func newJSONFileBackedConfigManager(file string) (Manager, error) {
	if file == "" {
		dir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		file = filepath.Join(dir, ".lan-tools.json")
	}
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("error reading configuration file: %w", err)
	}

	var c jsonConfig
	err = json.Unmarshal(data, &c)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling configuration file: %w", err)
	}

	m := &jsonFileBackedConfigManager{
		smtpConfig: &SMTPConfig{
			Host:   c.SMTPConfig.Host,
			Port:   c.SMTPConfig.Port,
			Sender: c.SMTPConfig.Sender,
			Pass:   c.SMTPConfig.Pass,
			To:     c.SMTPConfig.To,
		},
		databaseConfig: &DatabaseConfig{
			Path: c.DatabaseConfig.Path,
			Name: c.DatabaseConfig.Name,
		},
		cloudflareConfig: &CloudflareConfig{
			Token:   c.CloudflareConfig.Token,
			Domains: c.CloudflareConfig.Domains,
		},
		ispsConfig: make(map[string]*ISPConfig),
	}

	for k, v := range c.ISPsConfig {
		m.ispsConfig[k] = transform(v)
	}

	return m, nil
}

func transform(v jsonISPsConfig) *ISPConfig {
	return &ISPConfig{
		Name:               v.Name,
		PublicIPGetterType: v.PublicIPGetterType,
		PublicIPGetterCfg:  v.PublicIPGetterCfg,
		HostedDomains:      v.HostedDomains,
	}
}

func (cm *jsonFileBackedConfigManager) GetSMTPConfig() *SMTPConfig {
	return cm.smtpConfig
}

func (cm *jsonFileBackedConfigManager) GetDatabaseConfig() *DatabaseConfig {
	return cm.databaseConfig
}

func (cm *jsonFileBackedConfigManager) GetCloudflareConfig() *CloudflareConfig {
	return cm.cloudflareConfig
}

func (cm *jsonFileBackedConfigManager) GetISPsConfig() map[string]*ISPConfig {
	return cm.ispsConfig
}
