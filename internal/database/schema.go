package database

var ddls = []string{

	`CREATE TABLE IF NOT EXISTS isp (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		code TEXT UNIQUE NOT NULL,
		name TEXT UNIQUE NOT NULL,
		public_ip_getter_type TEXT NULL,
		public_ip_getter_cfg TEXT NULL,
		public_ip TEXT NULL,
		public_ip_modified TEXT NULL
	)`,

	`CREATE TABLE IF NOT EXISTS dns_provider (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		code TEXT UNIQUE NOT NULL,
		name TEXT UNIQUE NOT NULL,
		type TEXT NOT NULL,
		cfg TEXT NOT NULL
	)`,

	`CREATE TABLE IF NOT EXISTS domain (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT UNIQUE NOT NULL,
		description TEXT NULL,
        dns_provider_id INTEGER NOT NULL,
        FOREIGN KEY (dns_provider_id) REFERENCES dns_provider(id) 
	)`,

	`CREATE INDEX IF NOT EXISTS ix_for_domain_on_dns_provider_id ON domain(dns_provider_id)`,

	`CREATE TABLE IF NOT EXISTS dns_state (
		domain_id INTEGER,
		isp_id INTEGER,
		dns_provider_current_ip TEXT,
		dns_provider_record_id TEXT,
		dns_provider_sync_status TEXT NOT NULL,
		PRIMARY KEY (domain_id, isp_id),
		FOREIGN KEY (domain_id) REFERENCES domain(id),
		FOREIGN KEY (isp_id) REFERENCES isp(id)
	)`,

	`CREATE INDEX IF NOT EXISTS ix_for_dns_state_on_isp_id ON dns_state(isp_id)`,

	`CREATE TABLE IF NOT EXISTS isp_ip_change_log (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		code TEXT NOT NULL,
		public_ip_old TEXT NOT NULL,
		public_ip_new TEXT NOT NULL,
		created TEXT NOT NULL
	)`,

	`CREATE INDEX IF NOT EXISTS ix_for_isp_ip_change_log_on_code ON isp_ip_change_log(code)`,

	`CREATE TRIGGER trg_for_isp_after_update
	AFTER UPDATE ON isp
	FOR EACH ROW
	WHEN (OLD.public_ip IS NOT NEW.public_ip)
	BEGIN
	    INSERT INTO isp_ip_change_log(code, public_ip_old, public_ip_new, created)
	    VALUES (OLD.code, OLD.public_ip, NEW.public_ip, CURRENT_TIMESTAMP);
	END`,
}
