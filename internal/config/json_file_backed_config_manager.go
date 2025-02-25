package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

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

type jsonConfig struct {
	SMTPConfig     jsonSMTPConfig     `json:"smtp_config,omitempty"`
	DatabaseConfig jsonDatabaseConfig `json:"database_config,omitempty"`
}

type jsonFileBackedConfigManager struct {
	smtpConfig     *SMTPConfig
	databaseConfig *DatabaseConfig
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
	}

	return m, nil
}

func (cm *jsonFileBackedConfigManager) GetSMTPConfig() *SMTPConfig {
	return cm.smtpConfig
}

func (cm *jsonFileBackedConfigManager) GetDatabaseConfig() *DatabaseConfig {
	return cm.databaseConfig
}
