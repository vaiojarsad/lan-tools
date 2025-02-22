package isppublicip

import entities "github.com/vaiojarsad/cloudflare-tools/internal/entities/isppublicip"

type Dao interface {
	GetByISP(isp string) (*entities.ISPPublicIP, error)
	Insert(e *entities.ISPPublicIP) error
	UpdateIpInfo(e *entities.ISPPublicIP) error
}
