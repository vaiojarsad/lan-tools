package state

import (
	"fmt"

	"github.com/vaiojarsad/lan-tools/internal/dao"
	"github.com/vaiojarsad/lan-tools/internal/entities"
)

func Create(domainName, ispCode string) error {
	d, p, err := getKeyEntities(domainName, ispCode)
	if err != nil {
		return err
	}

	if d == nil {
		return fmt.Errorf("domain not found for name: %s", domainName)
	}

	if p == nil {
		return fmt.Errorf("isp not found for code: %s", ispCode)
	}

	domainISPCfgDao := dao.NewDnsStateDaoImpl()
	c, err := domainISPCfgDao.GetByDomainAndISPIds(d.StorageId(), p.StorageId())
	if err != nil {
		return fmt.Errorf("error searching domain isp configuration for domain %s and isp %s: %w", d.Name, p.Name, err)
	}

	if c != nil {
		return fmt.Errorf("configuration entry for domain %s and isp %s already exists. For updating, use 'refresh' instead of 'create'", d.Name, p.Name)
	}

	/*ip, err := isp.GetPublicIP(p.PublicIpGetterType, p.PublicIpGetterCfg)
	if err != nil {
		return fmt.Errorf("error retrieving public IP for ISP: %w", err)
	}

	err = isp.TryUpdateIspPublicIP(p, ip)
	if err != nil {
		environment.Get().ErrorLogger.Printf("error trying to update isp public IP in local DB: %v\n", err)
		// Proceed even if we fail to update locally
	}*/

	if err = domainISPCfgDao.Insert(&entities.DnsState{
		DomainId:              d.StorageId(),
		ISPId:                 p.StorageId(),
		DnsProviderCurrentIp:  entities.Unknown,
		DnsProviderRecordId:   entities.Unknown,
		DnsProviderSyncStatus: entities.Unknown, // Might have been synced externally.
	}); err != nil {
		return fmt.Errorf("error inserting domain-isp configuration: %w", err)
	}

	/*
		b, err := backend.NewDNSProviderBackendService(d.DnsProvider.ServiceType, d.DnsProvider.ServiceCfg)
		if err != nil {
			return fmt.Errorf("error getting backend service for %s: %w", d.DnsProvider.Name, err)
		}

		records, err := b.GetRecordsByTypeAndName(d.Name, "A", d.Name)
		if err != nil {
			return fmt.Errorf("error getting dns records from dns provider: %w", err)
		}
		environment.Get().OutputLogger.Println("", "records", records)
	*/

	return nil
}

func getKeyEntities(domainName, ispCode string) (*entities.Domain, *entities.ISP, error) {
	domainDao := dao.NewDomainDaoImpl(dao.NewDnsProviderDaoImpl())
	d, err := domainDao.GetByName(domainName)
	if err != nil {
		return nil, nil, fmt.Errorf("error searching domain by name: %w", err)
	}

	ispDao := dao.NewISPDaoImpl()
	p, err := ispDao.GetByCode(ispCode)
	if err != nil {
		return nil, nil, fmt.Errorf("error searching isp by code: %w", err)
	}

	return d, p, nil
}
