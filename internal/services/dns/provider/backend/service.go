package backend

import (
	"github.com/vaiojarsad/lan-tools/internal/services/dns/provider/backend/entities"
)

type Service interface {
	GetRecordsByTypeAndName(zone, rType, name string) ([]*entities.DNSRecord, error)
	UpdateDnsRecord(zone string, record *entities.DNSRecord) error
	CreateDnsRecord(zone string, record *entities.DNSRecord) error
	GetDnsRecord(zone, recordId string) (*entities.DNSRecord, error)
}
