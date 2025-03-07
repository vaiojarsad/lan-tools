package dao

import "github.com/vaiojarsad/lan-tools/internal/entities"

type DnsProviderDao interface {
	GetByCode(code string) (*entities.DnsProvider, error)
	GetById(id int64) (*entities.DnsProvider, error)
	Insert(e *entities.DnsProvider) error
}
