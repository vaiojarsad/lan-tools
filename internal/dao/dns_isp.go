package dao

import "github.com/vaiojarsad/lan-tools/internal/entities"

type DNSISPDao interface {
	GetByDomainAndISPIds(domainId, ispId int64) (*entities.DNSISP, error)
	GetByISPId(ispId int64) ([]*entities.DNSISP, error)
	Insert(e *entities.DNSISP) error
	UpdateDNSProviderInfo(e *entities.DNSISP) error
}
