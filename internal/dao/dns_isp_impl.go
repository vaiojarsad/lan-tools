package dao

import (
	"database/sql"
	"errors"

	"github.com/vaiojarsad/lan-tools/internal/database"
	"github.com/vaiojarsad/lan-tools/internal/entities"
	"github.com/vaiojarsad/lan-tools/internal/utils"
)

func NewDNSISPDaoImpl() DNSISPDao {
	return &databaseSqlDNSISPDaoImpl{}
}

type databaseSqlDNSISPDaoImpl struct {
}

func (d *databaseSqlDNSISPDaoImpl) Insert(e *entities.DNSISP) error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer utils.Close(db)

	stmt, err := db.Prepare("INSERT INTO dns_isp(domain_id, isp_id, dns_provider_record_id) VALUES (?, ?, ?)")
	if err != nil {
		return nil
	}
	defer utils.Close(stmt)

	_, err = stmt.Exec(e.DomainId, e.ISPId, e.DNSProviderRecordId)
	if err != nil {
		return err
	}

	return nil
}

func (d *databaseSqlDNSISPDaoImpl) GetByDomainAndISPIds(domainId, ispId int64) (*entities.DNSISP, error) {
	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer utils.Close(db)

	stmt, err := db.Prepare("SELECT dns_provider_record_id FROM dns_isp WHERE domain_id = ? and isp_id = ?")
	if err != nil {
		return nil, err
	}
	defer utils.Close(stmt)

	var dnsProviderRecordId string

	err = stmt.QueryRow(domainId, ispId).Scan(&dnsProviderRecordId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return entities.NewDNSISP(domainId, ispId, dnsProviderRecordId), nil
}

func (d *databaseSqlDNSISPDaoImpl) GetByISPId(ispId int64) ([]*entities.DNSISP, error) {
	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer utils.Close(db)

	stmt, err := db.Prepare("SELECT domain_id, dns_provider_record_id FROM dns_isp WHERE isp_id = ?")
	if err != nil {
		return nil, err
	}
	defer utils.Close(stmt)

	var domainId int64
	var dnsProviderRecordId string

	rows, err := stmt.Query(ispId)
	if err != nil {
		return nil, err
	}
	defer utils.Close(rows)

	var records []*entities.DNSISP
	for rows.Next() {
		err = rows.Scan(&domainId, &dnsProviderRecordId)
		if err != nil {
			return nil, err
		}
		records = append(records, entities.NewDNSISP(domainId, ispId, dnsProviderRecordId))
	}

	return records, nil
}

func (d *databaseSqlDNSISPDaoImpl) UpdateDNSProviderInfo(e *entities.DNSISP) error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer utils.Close(db)

	stmt, err := db.Prepare("UPDATE dns_isp SET dns_provider_record_id = ? WHERE domain_id = ? and isp_id = ?")
	if err != nil {
		return nil
	}
	defer utils.Close(stmt)

	_, err = stmt.Exec(e.DNSProviderRecordId, e.DomainId, e.ISPId)
	if err != nil {
		return err
	}

	return nil
}
