package dao

import (
	"database/sql"

	"github.com/vaiojarsad/lan-tools/internal/database"
	"github.com/vaiojarsad/lan-tools/internal/entities"
	"github.com/vaiojarsad/lan-tools/internal/utils"
)

func NewDnsStateDaoImpl() DnsStateDao {
	return &databaseSqlDnsStateDaoImpl{}
}

type databaseSqlDnsStateDaoImpl struct {
}

func (d *databaseSqlDnsStateDaoImpl) Insert(e *entities.DnsState) error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer utils.Close(db)

	stmt, err := db.Prepare("INSERT INTO dns_state(domain_id, isp_id, dns_provider_current_ip, dns_provider_record_id, dns_provider_sync_status) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return nil
	}
	defer utils.Close(stmt)

	_, err = stmt.Exec(e.DomainId, e.IspId, e.DnsProviderCurrentIp, e.DnsProviderRecordId, e.DnsProviderSyncStatus)
	if err != nil {
		return err
	}

	return nil
}

func (d *databaseSqlDnsStateDaoImpl) GetByDomainAndIspIds(domainId, ispId int64) (*entities.DnsState, error) {
	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer utils.Close(db)

	stmt, err := db.Prepare("SELECT dns_provider_current_ip, dns_provider_record_id, dns_provider_sync_status FROM dns_state WHERE domain_id = ? and isp_id = ?")
	if err != nil {
		return nil, err
	}
	defer utils.Close(stmt)

	var dnsProviderCurrentIp, dnsProviderRecordId, dnsProviderSyncStatus string

	err = stmt.QueryRow(domainId, ispId).Scan(&dnsProviderCurrentIp, &dnsProviderRecordId, &dnsProviderSyncStatus)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return entities.NewDnsState(domainId, ispId, dnsProviderCurrentIp, dnsProviderRecordId, dnsProviderSyncStatus), nil
}

func (d *databaseSqlDnsStateDaoImpl) UpdateDnsProviderInfo(e *entities.DnsState) error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer utils.Close(db)

	stmt, err := db.Prepare("UPDATE dns_state SET dns_provider_current_ip = ?, dns_provider_record_id = ?, dns_provider_sync_status = ? WHERE domain_id = ? and isp_id = ?")
	if err != nil {
		return nil
	}
	defer utils.Close(stmt)

	_, err = stmt.Exec(e.DnsProviderCurrentIp, e.DnsProviderRecordId, e.DnsProviderSyncStatus, e.DomainId, e.IspId)
	if err != nil {
		return err
	}

	return nil
}
