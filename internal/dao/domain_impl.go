package dao

import (
	"database/sql"
	"fmt"

	"github.com/vaiojarsad/lan-tools/internal/database"
	"github.com/vaiojarsad/lan-tools/internal/entities"
	"github.com/vaiojarsad/lan-tools/internal/utils"
)

func NewDomainDaoImpl(dnsProviderDao DNSProviderDao) DomainDao {
	return &databaseSqlDomainDaoImpl{dnsProviderDao: dnsProviderDao}
}

type databaseSqlDomainDaoImpl struct {
	dnsProviderDao DNSProviderDao
}

func (d *databaseSqlDomainDaoImpl) Insert(e *entities.Domain) error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer utils.Close(db)

	stmt, err := db.Prepare("INSERT INTO domain(name, description, dns_provider_id) VALUES (?, ?, ?)")
	if err != nil {
		return nil
	}
	defer utils.Close(stmt)

	_, err = stmt.Exec(e.Name, e.Description, e.DnsProviderStorageId())
	if err != nil {
		return err
	}

	return nil
}

func (d *databaseSqlDomainDaoImpl) GetByName(name string) (*entities.Domain, error) {
	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer utils.Close(db)

	stmt, err := db.Prepare("SELECT id, description, dns_provider_id FROM domain WHERE name = ?")
	if err != nil {
		return nil, err
	}
	defer utils.Close(stmt)

	var id, dnsProviderId int64
	var description string

	err = stmt.QueryRow(name).Scan(&id, &description, &dnsProviderId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	p, err := d.dnsProviderDao.GetById(dnsProviderId)
	if err != nil {
		return nil, fmt.Errorf("error retrieving dns provider: %w", err)
	}

	return entities.NewDomain(id, name, description, p), nil
}
