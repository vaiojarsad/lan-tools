package dao

import "github.com/vaiojarsad/lan-tools/internal/entities"

type ISPDao interface {
	GetByCode(code string) (*entities.ISP, error)
	Insert(e *entities.ISP) error
	UpdatePublicIpInfo(e *entities.ISP) error
}
