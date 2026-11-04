package backend

import (
	"github.com/vaiojarsad/lan-tools/internal/services/dns/provider/backend/entities"
)

type Service interface {
	GetRecordsByTypeAndName(zone, rType, name string) ([]*entities.DNSRecord, error)
	UpdateDNSRecord(zone string, record *entities.DNSRecord) error
	CreateDNSRecord(zone string, record *entities.DNSRecord) error
	GetDNSRecord(zone, recordId string) (*entities.DNSRecord, error)
}
