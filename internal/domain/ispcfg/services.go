package ispcfg

import (
	"fmt"

	"github.com/vaiojarsad/lan-tools/internal/dao"
)

func Create(domainName, ispCode string) error {
	domainDao := dao.NewDomainDaoImpl(dao.NewDNSProviderDaoImpl())
	d, err := domainDao.GetByName(domainName)
	if err != nil {
		return fmt.Errorf("error searching domain by name: %w", err)
	}

	if d == nil {
		return fmt.Errorf("domain not found for name: %s", domainName)
	}

	ispDao := dao.NewISPDaoImpl()
	p, err := ispDao.GetByCode(ispCode)
	if err != nil {
		return fmt.Errorf("error searching isp by code: %w", err)
	}

	if p == nil {
		return fmt.Errorf("isp not found for code: %s", ispCode)
	}

	return nil
}
