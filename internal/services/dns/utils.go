package dns

import (
	"github.com/vaiojarsad/lan-tools/internal/entities"
	dnsBackendEntities "github.com/vaiojarsad/lan-tools/internal/services/dns/provider/backend/entities"
)

// LookupRecordByISP that matches
func LookupRecordByISP(records []*dnsBackendEntities.DNSRecord, isp *entities.ISP) *dnsBackendEntities.DNSRecord {
	for _, record := range records {
		if isp.Code == record.ISPCode {
			return record
		}
	}
	return nil
}
