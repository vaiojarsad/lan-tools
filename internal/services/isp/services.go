package isp

import (
	"fmt"
	"time"

	"github.com/vaiojarsad/lan-tools/internal/dao"
	"github.com/vaiojarsad/lan-tools/internal/entities"
)

func Create(code, name, publicIpGetterType string, publicIpGetterCfg map[string]string) error {
	ispDao := dao.NewISPDaoImpl()

	isp := &entities.ISP{
		Code:               code,
		Name:               name,
		PublicIpGetterType: publicIpGetterType,
		PublicIpGetterCfg:  publicIpGetterCfg,
		PublicIp:           entities.Unknown,
		PublicIpModTime:    time.Time{},
	}
	if publicIpGetterType != "" {
		currentIp, err := GetPublicIP(publicIpGetterType, publicIpGetterCfg)
		if err != nil {
			return fmt.Errorf("error retrieving public IP for ISP: %w", err)
		}
		isp.PublicIp = currentIp
		isp.PublicIpModTime = time.Now()
	}
	if err := ispDao.Insert(isp); err != nil {
		return fmt.Errorf("error inserting ISP: %w", err)
	}
	return nil
}

func UpdateISPPublicIp(ispCode string) error {
	ispDao := dao.NewISPDaoImpl()

	isp, err := ispDao.GetByCode(ispCode)
	if err != nil {
		return fmt.Errorf("error looking up for ISP by code: %w", err)
	}

	if isp == nil {
		return fmt.Errorf("there is no ISP with code %s", ispCode)
	}

	ip, err := GetPublicIP(isp.PublicIpGetterType, isp.PublicIpGetterCfg)
	if err != nil {
		return fmt.Errorf("error retrieving public IP for ISP: %w", err)
	}

	if ip == isp.PublicIp {
		return nil
	}

	isp.PublicIp = ip
	isp.PublicIpModTime = time.Now()
	return ispDao.UpdatePublicIpInfo(isp)
}
