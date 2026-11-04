package dns

import (
	"errors"
	"fmt"

	"github.com/vaiojarsad/lan-tools/internal/dao"
	"github.com/vaiojarsad/lan-tools/internal/entities"
	"github.com/vaiojarsad/lan-tools/internal/environment"
	"github.com/vaiojarsad/lan-tools/internal/services/dns/provider/backend"
	backEntities "github.com/vaiojarsad/lan-tools/internal/services/dns/provider/backend/entities"
	ispServices "github.com/vaiojarsad/lan-tools/internal/services/isp"
)

func getDomainByName(domainName string) (*entities.Domain, error) {
	domainDao := dao.NewDomainDaoImpl(dao.NewDNSProviderDaoImpl())
	domain, err := domainDao.GetByName(domainName)
	if err != nil {
		return nil, err
	}

	if domain == nil {
		return nil, fmt.Errorf("domain not found for name: %s", domainName)
	}
	return domain, nil
}

func getISPByCode(ispCode string) (*entities.ISP, error) {
	ispDao := dao.NewISPDaoImpl()
	isp, err := ispDao.GetByCode(ispCode)
	if err != nil {
		return nil, err
	}

	if isp == nil {
		return nil, fmt.Errorf("isp not found for code: %s", ispCode)
	}
	return isp, nil
}

func getCurrentPublicIpForISP(isp *entities.ISP) (string, error) {
	ispCurrentIp, err := ispServices.GetPublicIP(isp.PublicIpGetterType, isp.PublicIpGetterCfg)
	if err != nil {
		return "", err
	}
	var localCurrentIp = isp.PublicIp
	if ispCurrentIp != localCurrentIp {
		environment.Get().OutputLogger.Printf("public ip (%s) from provider (%s) and local cache differ.\n", ispCurrentIp, isp.Name)
		err = ispServices.UpdateISPPublicIP(isp, ispCurrentIp)
		if err != nil {
			environment.Get().ErrorLogger.Printf("error trying to update isp public IP in local DB: %v\n", err)
		}
	} else {
		environment.Get().OutputLogger.Printf("public ip (%s) from provider (%s) and local cache (%s) match.\n", ispCurrentIp, isp.Name, isp.PublicIp)
	}
	return ispCurrentIp, nil
}

func updateOnIpChange(dnsISP *entities.DNSISP, isp *entities.ISP, domain *entities.Domain, newIP string) error {
	if dnsISP.DNSProviderRecordId == entities.Unknown {
		return fmt.Errorf("dns-isp relation for domain %s and isp %s doesn't have a valid provider record Id", domain.Name, isp.Name)
	}

	b, err := backend.NewDNSProviderBackendService(domain.DNSProvider.ServiceType, domain.DNSProvider.ServiceCfg)
	if err != nil {
		return fmt.Errorf("error getting backend service for %s: %w", domain.DNSProvider.Name, err)
	}

	record, err := b.GetDNSRecord(domain.Name, dnsISP.DNSProviderRecordId)
	if err != nil {
		return fmt.Errorf("error getting dns record from dns provider: %w", err)
	}

	if record.ISPCode != isp.Code || record.Content != newIP {
		record.ISPCode = isp.Code
		record.Content = newIP
		return b.UpdateDNSRecord(domain.Name, record)
	}

	return nil
}

func CreateRecordA(domainName, ispCode string) error {
	domain, err := getDomainByName(domainName)
	if err != nil {
		return fmt.Errorf("error searching domain by name: %w", err)
	}

	isp, err := getISPByCode(ispCode)
	if err != nil {
		return fmt.Errorf("error searching isp by code: %w", err)
	}

	ispCurrentIp, err := getCurrentPublicIpForISP(isp)
	if err != nil {
		return fmt.Errorf("error retrieving public IP for ISP: %w", err)
	}

	b, err := backend.NewDNSProviderBackendService(domain.DNSProvider.ServiceType, domain.DNSProvider.ServiceCfg)
	if err != nil {
		return fmt.Errorf("error getting backend service for %s: %w", domain.DNSProvider.Name, err)
	}

	records, err := b.GetRecordsByTypeAndName(domain.Name, "A", domain.Name)
	if err != nil {
		return fmt.Errorf("error getting dns records from dns provider: %w", err)
	}

	record := LookupRecordByISP(records, isp)
	if record == nil {
		record = &backEntities.DNSRecord{
			Type:    "A",
			Name:    domainName,
			Content: ispCurrentIp,
			ISPCode: isp.Code,
		}
		err = b.CreateDNSRecord(domain.Name, record)
	} else {
		if record.ISPCode != isp.Code || record.Content != ispCurrentIp {
			record.ISPCode = isp.Code
			record.Content = ispCurrentIp
			err = b.UpdateDNSRecord(domain.Name, record)
		}
	}
	if err != nil {
		return err
	}

	dnsISPDao := dao.NewDNSISPDaoImpl()
	dnsISP, err := dnsISPDao.GetByDomainAndISPIds(domain.StorageId(), isp.StorageId())
	if err != nil {
		return err
	}

	if dnsISP != nil {
		if dnsISP.DNSProviderRecordId != record.ProviderId {
			dnsISP.DNSProviderRecordId = record.ProviderId
			err = dnsISPDao.UpdateDNSProviderInfo(dnsISP)
		}
	} else {
		err = dnsISPDao.Insert(&entities.DNSISP{
			DomainId:            domain.StorageId(),
			ISPId:               isp.StorageId(),
			DNSProviderRecordId: record.ProviderId,
		})
	}

	return err
}

func SyncRecordA(domainName, ispCode string) error {
	domain, err := getDomainByName(domainName)
	if err != nil {
		return fmt.Errorf("error searching domain by name: %w", err)
	}

	isp, err := getISPByCode(ispCode)
	if err != nil {
		return fmt.Errorf("error searching isp by code: %w", err)
	}

	ispCurrentIp, err := getCurrentPublicIpForISP(isp)
	if err != nil {
		return fmt.Errorf("error retrieving public IP for ISP: %w", err)
	}

	dnsISPDao := dao.NewDNSISPDaoImpl()
	dnsISP, err := dnsISPDao.GetByDomainAndISPIds(domain.StorageId(), isp.StorageId())
	if err != nil {
		return fmt.Errorf("error searching dns-isp relation for domain %s and isp %s: %w", domain.Name, isp.Name, err)
	}

	if dnsISP == nil {
		return fmt.Errorf("dns-isp relation entry for domain %s and isp %s wasn't found", domain.Name, isp.Name)
	}

	return updateOnIpChange(dnsISP, isp, domain, ispCurrentIp)
}

func SyncRecordsA(ispCode string) error {
	isp, err := getISPByCode(ispCode)
	if err != nil {
		return fmt.Errorf("error searching isp by code: %w", err)
	}

	ispCurrentIp, err := getCurrentPublicIpForISP(isp)
	if err != nil {
		return fmt.Errorf("error retrieving public IP for ISP: %w", err)
	}

	if err = sendIpChangedMail(isp.Name, ispCurrentIp, isp.PublicIp); err != nil {
		environment.Get().ErrorLogger.Printf("error sending ip changed mail: %v\n", err)
	}

	dnsISPDao := dao.NewDNSISPDaoImpl()
	dnsISPs, err := dnsISPDao.GetByISPId(isp.StorageId())
	if err != nil {
		return fmt.Errorf("error searching dns-isp relations for isp %s: %w", isp.Name, err)
	}

	var errs []error
	domainDao := dao.NewDomainDaoImpl(dao.NewDNSProviderDaoImpl())

	for _, dnsISP := range dnsISPs {
		domain, err2 := domainDao.GetById(dnsISP.DomainId)
		if err2 != nil {
			errs = append(errs, fmt.Errorf("error searching domain by id (%d): %w", dnsISP.DomainId, err2))
			continue
		}

		err2 = updateOnIpChange(dnsISP, isp, domain, ispCurrentIp)
		if err2 != nil {
			errs = append(errs, fmt.Errorf("error trying to update ip for domain %s and isp %s: %w", domain.Name, isp.Name, err2))
			continue
		}
	}

	if len(errs) > 0 {
		if err = sendSyncErrorsMail(isp.Name, ispCurrentIp, errs); err != nil {
			environment.Get().ErrorLogger.Printf("error sending sync errors mail: %v\n", err)
		}
	}

	return errors.Join(errs...)
}
