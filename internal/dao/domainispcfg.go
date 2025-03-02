package dao

import "github.com/vaiojarsad/lan-tools/internal/entities"

type DomainISPCfgDao interface {
	GetByDomainAndISPIds(domainId, ispId int64) (*entities.DomainISPCfg, error)
	Insert(e *entities.DomainISPCfg) error
	UpdateDnsProviderInfo(e *entities.DomainISPCfg) error
}
