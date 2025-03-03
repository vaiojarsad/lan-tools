package dao

import (
	"database/sql"
	"encoding/json"
	"github.com/vaiojarsad/lan-tools/internal/database"
	"github.com/vaiojarsad/lan-tools/internal/entities"
	"github.com/vaiojarsad/lan-tools/internal/utils"
)

func NewDNSProviderDaoImpl() DNSProviderDao {
	return &databaseSqlDnsProviderDaoImpl{}
}

type databaseSqlDnsProviderDaoImpl struct {
}

func (d *databaseSqlDnsProviderDaoImpl) Insert(e *entities.DNSProvider) error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer utils.Close(db)

	stmt, err := db.Prepare("INSERT INTO dns_provider(code, name, type, cfg) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil
	}
	defer utils.Close(stmt)

	cfg, err := json.Marshal(e.ServiceCfg)
	if err != nil {
		return nil
	}

	_, err = stmt.Exec(e.Code, e.Name, e.ServiceType, cfg)
	if err != nil {
		return err
	}

	return nil
}

func (d *databaseSqlDnsProviderDaoImpl) GetByCode(code string) (*entities.DNSProvider, error) {
	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer utils.Close(db)

	stmt, err := db.Prepare("SELECT id, name, type, cfg FROM dns_provider WHERE code = ?")
	if err != nil {
		return nil, err
	}
	defer utils.Close(stmt)

	var id int64
	var name, serviceType, serviceCfgStr string

	err = stmt.QueryRow(code).Scan(&id, &name, &serviceType, &serviceCfgStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	var serviceCfg map[string]string
	err = json.Unmarshal([]byte(serviceCfgStr), &serviceCfg)
	if err != nil {
		return nil, err
	}

	return entities.NewDNSProvider(id, code, name, serviceType, serviceCfg), nil
}

func (d *databaseSqlDnsProviderDaoImpl) GetById(id int64) (*entities.DNSProvider, error) {
	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer utils.Close(db)

	stmt, err := db.Prepare("SELECT code, name, type, cfg FROM dns_provider WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer utils.Close(stmt)

	var code, name, serviceType, serviceCfgStr string

	err = stmt.QueryRow(id).Scan(&code, &name, &serviceType, &serviceCfgStr)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	var serviceCfg map[string]string
	err = json.Unmarshal([]byte(serviceCfgStr), &serviceCfg)
	if err != nil {
		return nil, err
	}

	return entities.NewDNSProvider(id, code, name, serviceType, serviceCfg), nil
}
