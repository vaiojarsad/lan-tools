package dao

import "github.com/vaiojarsad/lan-tools/internal/entities"

type DNSProviderDao interface {
	GetByCode(code string) (*entities.DNSProvider, error)
	Insert(e *entities.DNSProvider) error
}
