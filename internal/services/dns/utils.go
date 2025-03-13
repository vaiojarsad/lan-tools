package dns

import (
	"errors"
	"github.com/vaiojarsad/lan-tools/internal/entities"
	dnsBackendEntities "github.com/vaiojarsad/lan-tools/internal/services/dns/provider/backend/entities"
)

var ErrIspCodeNotSet = errors.New("isp code is not set on provider record")
var ErrIspMismatch = errors.New("record ownership is set to an unexpected isp on provider record")

type lookupRecord func(records []*dnsBackendEntities.DNSRecord, dnsState *entities.DnsState, isp *entities.Isp, currentProviderIp, currentLocalIp string) (*dnsBackendEntities.DNSRecord, error)

func WrapWithSanityCheck(fn lookupRecord) lookupRecord {
	return func(records []*dnsBackendEntities.DNSRecord, dnsState *entities.DnsState, isp *entities.Isp, currentProviderIp, currentLocalIp string) (*dnsBackendEntities.DNSRecord, error) {
		ispRecordsCount := 0
		var recordForCurrentProviderIp *dnsBackendEntities.DNSRecord = nil
		var recordForCurrentLocalIp *dnsBackendEntities.DNSRecord = nil
		for _, record := range records {
			if isp.Code == record.IspCode {
				ispRecordsCount++
			}
			if currentProviderIp == record.Content {
				recordForCurrentProviderIp = record
			}
			if currentProviderIp == record.Content {
				recordForCurrentLocalIp = record
			}
		}

		var errs []error
		if ispRecordsCount > 1 {
			errs = append(errs, errors.New("more than one record marked with the isp code were found on dns provider side"))
		}
		if recordForCurrentProviderIp != nil && recordForCurrentProviderIp.IspCode == "" && ispRecordsCount > 0 {
			errs = append(errs, errors.New("unowned record with the current public ip for the isp and at least one other record marked as owned by the isp were found on provider side"))
		}
		if recordForCurrentLocalIp != nil && recordForCurrentLocalIp.IspCode == "" && ispRecordsCount > 0 {
			errs = append(errs, errors.New("unowned record with the current local public ip for the isp and at least one other record marked as owned by the isp were found on provider side"))
		}

		if len(errs) > 0 {
			return nil, errors.Join(errs...)
		}

		return fn(records, dnsState, isp, currentProviderIp, currentLocalIp)
	}
}

// LookupRecord that matches
func LookupRecord(records []*dnsBackendEntities.DNSRecord, dnsState *entities.DnsState, isp *entities.Isp, currentProviderIp, currentLocalIp string) (
	*dnsBackendEntities.DNSRecord, error) {
	for _, record := range records {
		if dnsState != nil {
			if dnsState.DnsProviderRecordId == record.ProviderId {
				if isp.Code == record.IspCode {
					return record, nil
				}
				if record.IspCode == "" {
					return record, ErrIspCodeNotSet
				}
				return record, ErrIspMismatch
			}
		}
		if currentProviderIp == record.Content {
			if isp.Code == record.IspCode {
				return record, nil
			}
			if record.IspCode == "" {
				return record, ErrIspCodeNotSet
			}
			return record, ErrIspMismatch
		}
		if currentLocalIp == record.Content {
			if isp.Code == record.IspCode {
				return record, nil
			}
			if record.IspCode == "" {
				return record, ErrIspCodeNotSet
			}
			return record, ErrIspMismatch
		}
		if isp.Code == record.IspCode {
			return record, nil
		}
	}
	return nil, nil
}
