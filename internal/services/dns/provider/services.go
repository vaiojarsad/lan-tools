package provider

import (
	"github.com/vaiojarsad/lan-tools/internal/dao"
	"github.com/vaiojarsad/lan-tools/internal/entities"
)

func Create(code, name, serviceType string, serviceCfg map[string]string) error {
	dnsProviderDao := dao.NewDNSProviderDaoImpl()
	dnsProvider := &entities.DNSProvider{
		Code:        code,
		Name:        name,
		ServiceType: serviceType,
		ServiceCfg:  serviceCfg,
	}

	return dnsProviderDao.Insert(dnsProvider)
}
