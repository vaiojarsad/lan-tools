package database

var tables = []string{
	`CREATE TABLE IF NOT EXISTS isp_public_ip (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		isp TEXT UNIQUE NOT NULL,
		ip TEXT UNIQUE NOT NULL,
		modified TEXT NOT NULL
	)`,
}
