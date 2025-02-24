package database

import (
	"database/sql"
	"path"

	_ "modernc.org/sqlite"

	"github.com/vaiojarsad/lan-tools/internal/environment"
)

func getDatabasePath() string {
	cfg := environment.Get().ConfigManager.GetDatabaseConfig()
	return path.Join(cfg.Path, cfg.Name)
}

func Open() (*sql.DB, error) {
	return sql.Open("sqlite", getDatabasePath())
}
