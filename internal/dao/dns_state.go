package dao

import "github.com/vaiojarsad/lan-tools/internal/entities"

type DnsStateDao interface {
	GetByDomainAndIspIds(domainId, ispId int64) (*entities.DnsState, error)
	Insert(e *entities.DnsState) error
	UpdateDnsProviderInfo(e *entities.DnsState) error
}
