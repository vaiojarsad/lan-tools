package config

import (
	"sync"
)

var (
	instance Manager
	once     sync.Once
)

// Manager handle the different configurations
type Manager interface {
	GetSMTPConfig() *SMTPConfig
	GetDatabaseConfig() *DatabaseConfig
}

// Create returns the singleton instance, creating it if necessary
func Create(configFile string) (Manager, error) {
	var err error = nil
	once.Do(func() {
		instance, err = newJSONFileBackedConfigManager(configFile)
	})
	return instance, err
}

func Get() Manager {
	return instance
}
