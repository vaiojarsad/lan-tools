package dao

import "github.com/vaiojarsad/lan-tools/internal/entities"

type DomainDao interface {
	GetByName(name string) (*entities.Domain, error)
	Insert(e *entities.Domain) error
}
