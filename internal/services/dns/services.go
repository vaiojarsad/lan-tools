package dns

import (
	"errors"
	"fmt"
	backEntities "github.com/vaiojarsad/lan-tools/internal/services/dns/provider/backend/entities"

	"github.com/vaiojarsad/lan-tools/internal/dao"
	"github.com/vaiojarsad/lan-tools/internal/entities"
	"github.com/vaiojarsad/lan-tools/internal/environment"
	"github.com/vaiojarsad/lan-tools/internal/services/dns/provider/backend"
	stateServices "github.com/vaiojarsad/lan-tools/internal/services/dns/state"
	ispServices "github.com/vaiojarsad/lan-tools/internal/services/isp"
)

func CreateRecordA(domainName, ispCode string) error {
	domain, isp, err := stateServices.GetKeyEntities(domainName, ispCode)
	if err != nil {
		return err
	}

	ispCurrentIp, err := ispServices.GetPublicIP(isp.PublicIpGetterType, isp.PublicIpGetterCfg)
	if err != nil {
		return fmt.Errorf("error retrieving public IP for ISP: %w", err)
	}

	var localCurrentIp = isp.PublicIp
	if ispCurrentIp != localCurrentIp {
		err = ispServices.UpdateIspPublicIP(isp, ispCurrentIp)
		if err != nil {
			environment.Get().ErrorLogger.Printf("error trying to update isp public IP in local DB: %v\n", err)
			// Proceed even if we fail to update locally
		}
	}

	dnsStateDao := dao.NewDnsStateDaoImpl()
	state, err := dnsStateDao.GetByDomainAndIspIds(domain.StorageId(), isp.StorageId())
	if err != nil {
		return fmt.Errorf("error searching domain isp configuration for domain %state and isp %state: %w", domain.Name, isp.Name, err)
	}

	if state != nil {
		// The ip we got from isp.GetPublicIP matches the one we have in the local DB as the last IP we informed about to DNS provider.
		// Strictly speaking the IP might have been externally altered in the dns provider. By default, we assume it wasn't the case.
		if state.DnsProviderCurrentIp == ispCurrentIp {
			return nil
		}
	}

	b, err := backend.NewDNSProviderBackendService(domain.DnsProvider.ServiceType, domain.DnsProvider.ServiceCfg)
	if err != nil {
		return fmt.Errorf("error getting backend service for %s: %w", domain.DnsProvider.Name, err)
	}

	records, err := b.GetRecordsByTypeAndName(domain.Name, "A", domain.Name)
	if err != nil {
		return fmt.Errorf("error getting dns records from dns provider: %w", err)
	}

	record, err := WrapWithSanityCheck(LookupRecord)(records, state, isp, ispCurrentIp, localCurrentIp)
	if err != nil {
		if errors.Is(err, ErrIspMismatch) {
			return fmt.Errorf("inconsistency error: %w", err)
		}
		if errors.Is(err, ErrIspCodeNotSet) {
			environment.Get().OutputLogger.Printf("unowned dns provider record (id %s). Will be updated. \n", record.ProviderId)
		} else {
			return err
		}
	}

	if record == nil {
		record = &backEntities.DNSRecord{
			Type:    "A",
			Name:    domainName,
			Content: ispCurrentIp,
			IspCode: isp.Code,
		}
		err = b.CreateDnsRecord(domain.Name, record)
	} else {
		if record.IspCode != isp.Code || record.Content != ispCurrentIp {
			record.IspCode = isp.Code
			record.Content = ispCurrentIp
			err = b.UpdateDnsRecord(domain.Name, record)
		}
	}
	if err != nil {
		return err
	}

	if state != nil {
		state.DnsProviderRecordId = record.ProviderId
		state.DnsProviderSyncStatus = entities.Synced
		state.DnsProviderCurrentIp = ispCurrentIp
		err = dnsStateDao.UpdateDnsProviderInfo(state)
	} else {
		err = dnsStateDao.Insert(&entities.DnsState{
			DomainId:              domain.StorageId(),
			IspId:                 isp.StorageId(),
			DnsProviderCurrentIp:  ispCurrentIp,
			DnsProviderRecordId:   record.ProviderId,
			DnsProviderSyncStatus: entities.Synced,
		})
	}
	return err
}
