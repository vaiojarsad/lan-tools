package state

import (
	"fmt"

	"github.com/vaiojarsad/lan-tools/internal/dao"
	"github.com/vaiojarsad/lan-tools/internal/entities"
)

func Create(domainName, ispCode string) error {
	d, p, err := GetKeyEntities(domainName, ispCode)
	if err != nil {
		return err
	}

	dnsStateDao := dao.NewDnsStateDaoImpl()
	s, err := dnsStateDao.GetByDomainAndIspIds(d.StorageId(), p.StorageId())
	if err != nil {
		return fmt.Errorf("error searching domain isp configuration for domain %s and isp %s: %w", d.Name, p.Name, err)
	}

	if s != nil {
		return fmt.Errorf("configuration entry for domain %s and isp %s already exists. For updating, use 'refresh' instead of 'create'", d.Name, p.Name)
	}

	if err = dnsStateDao.Insert(&entities.DnsState{
		DomainId:              d.StorageId(),
		IspId:                 p.StorageId(),
		DnsProviderCurrentIp:  entities.Unknown,
		DnsProviderRecordId:   entities.Unknown,
		DnsProviderSyncStatus: entities.Unknown, // Might have been synced externally.
	}); err != nil {
		return fmt.Errorf("error inserting domain-isp configuration: %w", err)
	}

	return nil
}

func GetKeyEntities(domainName, ispCode string) (*entities.Domain, *entities.Isp, error) {
	domainDao := dao.NewDomainDaoImpl(dao.NewDnsProviderDaoImpl())
	d, err := domainDao.GetByName(domainName)
	if err != nil {
		return nil, nil, fmt.Errorf("error searching domain by name: %w", err)
	}

	if d == nil {
		return nil, nil, fmt.Errorf("domain not found for name: %s", domainName)
	}

	ispDao := dao.NewISPDaoImpl()
	p, err := ispDao.GetByCode(ispCode)
	if err != nil {
		return nil, nil, fmt.Errorf("error searching isp by code: %w", err)
	}

	if p == nil {
		return nil, nil, fmt.Errorf("isp not found for code: %s", ispCode)
	}

	return d, p, nil
}
