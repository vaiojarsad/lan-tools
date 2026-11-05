package isp

import (
	"fmt"
	"time"

	"github.com/vaiojarsad/lan-tools/internal/dao"
	"github.com/vaiojarsad/lan-tools/internal/entities"
	"github.com/vaiojarsad/lan-tools/internal/environment"
	"github.com/vaiojarsad/lan-tools/internal/services/mail"
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

func CheckISPPublicIp(ispCode string, trigger bool) error {
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

	if err = mail.SendIpChangedMail(isp.Name, ip, isp.PublicIp); err != nil {
		environment.Get().ErrorLogger.Printf("error sending ip changed mail: %v\n", err)
	}

	if ip == isp.PublicIp || trigger == false {
		return nil
	}
	
	isp.PublicIp = ip
	isp.PublicIpModTime = time.Now()
	return ispDao.UpdatePublicIpInfo(isp)
}

func GetISPByCode(ispCode string) (*entities.ISP, error) {
	ispDao := dao.NewISPDaoImpl()
	isp, err := ispDao.GetByCode(ispCode)
	if err != nil {
		return nil, err
	}

	if isp == nil {
		return nil, fmt.Errorf("isp not found for code: %s", ispCode)
	}
	return isp, nil
}
