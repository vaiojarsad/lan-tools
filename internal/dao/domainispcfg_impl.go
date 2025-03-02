package dao

import (
	"database/sql"

	"github.com/vaiojarsad/lan-tools/internal/database"
	"github.com/vaiojarsad/lan-tools/internal/entities"
	"github.com/vaiojarsad/lan-tools/internal/utils"
)

func NewDomainISPCfgDaoImpl() DomainISPCfgDao {
	return &databaseSqlDomainIspCfgDaoImpl{}
}

type databaseSqlDomainIspCfgDaoImpl struct {
}

func (d *databaseSqlDomainIspCfgDaoImpl) Insert(e *entities.DomainISPCfg) error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer utils.Close(db)

	stmt, err := db.Prepare("INSERT INTO domain_isp_cfg(domain_id, isp_id, dns_provider_current_ip, dns_provider_record_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil
	}
	defer utils.Close(stmt)

	_, err = stmt.Exec(e.DomainId, e.ISPId, e.DnsProviderCurrentIp, e.DnsProviderRecordId)
	if err != nil {
		return err
	}

	return nil
}

func (d *databaseSqlDomainIspCfgDaoImpl) GetByDomainAndISPIds(domainId, ispId int64) (*entities.DomainISPCfg, error) {
	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer utils.Close(db)

	stmt, err := db.Prepare("SELECT dns_provider_current_ip, dns_provider_record_id FROM domain_isp_cfg WHERE domain_id = ? and isp_id = ?")
	if err != nil {
		return nil, err
	}
	defer utils.Close(stmt)

	var dnsProviderCurrentIp, dnsProviderRecordId string

	err = stmt.QueryRow(domainId, ispId).Scan(&dnsProviderCurrentIp, &dnsProviderRecordId)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return entities.NewDomainISPCfg(domainId, ispId, dnsProviderCurrentIp, dnsProviderRecordId), nil
}

func (d *databaseSqlDomainIspCfgDaoImpl) UpdateDnsProviderInfo(e *entities.DomainISPCfg) error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer utils.Close(db)

	stmt, err := db.Prepare("UPDATE domain_isp_cfg SET dns_provider_current_ip = ?, dns_provider_record_id = ? WHERE domain_id = ? and isp_id = ?")
	if err != nil {
		return nil
	}
	defer utils.Close(stmt)

	_, err = stmt.Exec(e.DnsProviderCurrentIp, e.DnsProviderRecordId, e.DomainId, e.ISPId)
	if err != nil {
		return err
	}

	return nil
}
