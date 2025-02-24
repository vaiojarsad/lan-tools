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
        dns_provider_id INTEGER UNIQUE NOT NULL,
        FOREIGN KEY (dns_provider_id) REFERENCES dns_provider(id) 
	)`,

	`CREATE TABLE IF NOT EXISTS domain_isp_cfg (
		domain_id INTEGER,
		isp_id INTEGER,
		dns_provider_current_ip TEXT NOT NULL,
		dns_provider_record_id TEXT NOT NULL,
		PRIMARY KEY (domain_id, isp_id),
		FOREIGN KEY (domain_id) REFERENCES domain(id),
		FOREIGN KEY (isp_id) REFERENCES isp(id),
        CONSTRAINT uk_domain_isp UNIQUE (domain_id, dns_provider_current_ip),
        CONSTRAINT uk_domain_isp UNIQUE (domain_id, dns_provider_record_id)
	)`,

	`CREATE INDEX IF NOT EXISTS ix_domain_isp_cfg_isp_id ON domain_isp_cfg(isp_id)`,
}

/*


{
        "smtp_config": {
            "host": "smtp.gmail.com",
            "port": 587,
            "sender": "daniel.huespe.oso@gmail.com",
            "password": "vbwfkmyizsfdbhbw",
            "to": "pdcvgmh@gmail.com"
        },
		"isps_config": {
			"mov": {
				"name": "Movistar",
				"public_ip_getter_type": "ipify",
				"public_ip_getter_cfg": {
					"url": "https://api64.ipify.org",
					"ip": "64.185.227.155"
				},
				"hosted_domains": [
					"19741976.xyz",
					"riggedsystems.us"
				]
			},
			"tel": {
				"name": "Telecentro",
				"public_ip_getter_type": "ipify",
				"public_ip_getter_cfg": {
					"url": "https://api64.ipify.org",
					"ip": "173.231.16.77"
				},
				"hosted_domains": [
					"19741976.xyz",
					"riggedsystems.us"
				]
			}
		},
		"database_config": {
			"path": "",
			"name": "lan-tools-local.db"
		},
		"cloudflare_config": {
			"token": "mCMMGvgRBbhj5uoD57ia0t0OYK2BJ3Hm29eFBbLO",
			"domains": [
				"19741976.xyz",
				"riggedsystems.us"
			]
		}
}



*/
