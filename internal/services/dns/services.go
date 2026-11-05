package dns

import (
	"errors"
	"fmt"

	"github.com/vaiojarsad/lan-tools/internal/dao"
	"github.com/vaiojarsad/lan-tools/internal/entities"
	"github.com/vaiojarsad/lan-tools/internal/environment"
	"github.com/vaiojarsad/lan-tools/internal/services/dns/provider/backend"
	backEntities "github.com/vaiojarsad/lan-tools/internal/services/dns/provider/backend/entities"
	domainServices "github.com/vaiojarsad/lan-tools/internal/services/domain"
	ispServices "github.com/vaiojarsad/lan-tools/internal/services/isp"
	"github.com/vaiojarsad/lan-tools/internal/services/mail"
)

func getCommons(domainName, ispCode string) (*entities.ISP, *entities.Domain, error) {
	domain, err := domainServices.GetDomainByName(domainName)
	if err != nil {
		return nil, nil, fmt.Errorf("error searching domain by name: %w", err)
	}

	isp, err := ispServices.GetISPByCode(ispCode)
	if err != nil {
		return nil, nil, fmt.Errorf("error searching isp by code: %w", err)
	}
	return isp, domain, nil
}


func CreateRecordA(domainName, ispCode string) error {
	isp, domain, err := getCommons(domainName, ispCode)
	if err != nil {
		return err
	}

	ispCurrentIp, err := ispServices.GetPublicIP(isp.PublicIpGetterType, isp.PublicIpGetterCfg)
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
	isp, domain, err := getCommons(domainName, ispCode)
	if err != nil {
		return err
	}

	ispCurrentIp, err := ispServices.GetPublicIP(isp.PublicIpGetterType, isp.PublicIpGetterCfg)
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

	return updateRecordA(dnsISP, isp, domain, ispCurrentIp)
}

func SyncRecordsA(ispCode string) error {
	isp, err := ispServices.GetISPByCode(ispCode)
	if err != nil {
		return fmt.Errorf("error searching isp by code: %w", err)
	}

	ispCurrentIp, err := ispServices.GetPublicIP(isp.PublicIpGetterType, isp.PublicIpGetterCfg)
	if err != nil {
		return fmt.Errorf("error retrieving public IP for ISP: %w", err)
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

		err2 = updateRecordA(dnsISP, isp, domain, ispCurrentIp)
		if err2 != nil {
			errs = append(errs, fmt.Errorf("error trying to update ip for domain %s and isp %s: %w", domain.Name, isp.Name, err2))
			continue
		}
	}

	if len(errs) > 0 {
		if err = mail.SendSyncErrorsMail(isp.Name, ispCurrentIp, errs); err != nil {
			environment.Get().ErrorLogger.Printf("error sending sync errors mail: %v\n", err)
		}
	}

	return errors.Join(errs...)
}

func updateRecordA(dnsISP *entities.DNSISP, isp *entities.ISP, domain *entities.Domain, newIP string) error {
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
