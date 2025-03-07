package provider

import (
	"fmt"

	"github.com/vaiojarsad/lan-tools/internal/dao"
	"github.com/vaiojarsad/lan-tools/internal/entities"
)

func Create(code, name, serviceType string, serviceCfg map[string]string) error {
	dnsProviderDao := dao.NewDnsProviderDaoImpl()
	dnsProvider := &entities.DnsProvider{
		Code:        code,
		Name:        name,
		ServiceType: serviceType,
		ServiceCfg:  serviceCfg,
	}

	if err := dnsProviderDao.Insert(dnsProvider); err != nil {
		return fmt.Errorf("error inserting dns provider: %w", err)
	}
	return nil
}
