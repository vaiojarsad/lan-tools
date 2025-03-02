package domain

import (
	"fmt"

	"github.com/vaiojarsad/lan-tools/internal/dao"
	"github.com/vaiojarsad/lan-tools/internal/entities"
)

func Create(name, description, dnsProviderCode string) error {
	dnsProviderDao := dao.NewDNSProviderDaoImpl()

	dnsProvider, err := dnsProviderDao.GetByCode(dnsProviderCode)
	if err != nil {
		return fmt.Errorf("error searching dns provider by code: %w", err)
	}

	if dnsProvider == nil {
		return fmt.Errorf("dns provider not found for code: %s", dnsProviderCode)
	}

	domainDao := dao.NewDomainDaoImpl(dnsProviderDao)

	domain := &entities.Domain{
		Name:        name,
		Description: description,
		DnsProvider: dnsProvider,
	}

	return domainDao.Insert(domain)
}
