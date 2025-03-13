package dao

import "github.com/vaiojarsad/lan-tools/internal/entities"

type ISPDao interface {
	GetByCode(code string) (*entities.Isp, error)
	Insert(e *entities.Isp) error
	UpdatePublicIpInfo(e *entities.Isp) error
}
