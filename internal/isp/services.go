package isp

import (
	"fmt"
	"time"

	"github.com/vaiojarsad/lan-tools/internal/dao"
	"github.com/vaiojarsad/lan-tools/internal/entities"
	"github.com/vaiojarsad/lan-tools/internal/environment"
)

func Create(code, name, publicIpGetterType string, publicIpGetterCfg map[string]string) error {
	ispDao := dao.NewISPDaoImpl()
	isp := &entities.ISP{
		Code:               code,
		Name:               name,
		PublicIpGetterType: publicIpGetterType,
		PublicIpGetterCfg:  publicIpGetterCfg,
		PublicIp:           "",
		PublicIpModTime:    time.Time{},
	}
	if publicIpGetterType != "" {
		currentIp, err := GetPublicIP(publicIpGetterType, publicIpGetterCfg)
		if err != nil {
			return err
		}
		isp.PublicIp = currentIp
		isp.PublicIpModTime = time.Now()
	}
	return ispDao.Insert(isp)
}

func RefreshIspPublicIp(ispCode string) error {
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

	return TryUpdateIspPublicIP(isp, ip)
}

func TryUpdateIspPublicIP(isp *entities.ISP, ip string) error {
	ispDao := dao.NewISPDaoImpl()
	if ip != isp.PublicIp {
		environment.Get().OutputLogger.Printf("updating public ip for %s. Old: %s New: %s\n", isp.Name, isp.PublicIp, ip)
		isp.PublicIp = ip
		isp.PublicIpModTime = time.Now()
		if err := ispDao.UpdatePublicIpInfo(isp); err != nil {
			return fmt.Errorf("error updating public ip info: %w", err)
		}
	} else {
		environment.Get().OutputLogger.Printf("local and current public ip values for %s are equal. %s\n", isp.Name, ip)
	}
	return nil
}
