package environment

import (
	"log"
	"sync"

	"github.com/vaiojarsad/lan-tools/internal/config"
)

// Environment holds the execution context
type Environment struct {
	ConfigManager config.Manager
	ErrorLogger   *log.Logger
	OutputLogger  *log.Logger
}

var (
	instance *Environment
	once     sync.Once
)

// Create returns the singleton instance, creating it if necessary
func Create() *Environment {
	once.Do(func() {
		instance = &Environment{}
	})
	return instance
}

func Get() *Environment {
	return instance
}
