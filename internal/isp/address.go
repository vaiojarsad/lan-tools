package isp

import (
	"errors"
	"fmt"
	"time"

	"github.com/vaiojarsad/cloudflare-tools/internal/config"
	"github.com/vaiojarsad/cloudflare-tools/internal/dao/isppublicip"
	entities "github.com/vaiojarsad/cloudflare-tools/internal/entities/isppublicip"
	"github.com/vaiojarsad/cloudflare-tools/internal/environment"
	"github.com/vaiojarsad/cloudflare-tools/internal/public_ip_resolver"
)

func GetPublicIP(cfg *config.ISPConfig) (string, error) {
	r, err := public_ip_resolver.NewPublicIPResolver(cfg.PublicIPGetterType, cfg.PublicIPGetterCfg)
	if err != nil {
		return "", err
	}
	ip, err := r.Resolve()
	if err != nil {
		return "", err
	}
	return ip, nil
}

type ispCfgWithId struct {
	ispId  string
	ispCfg *config.ISPConfig
}

func UpdatePublicIP(ispsIds []string) error {
	var errs []error
	var ispsToUpdate []*ispCfgWithId
	ispsCfg := environment.Get().ConfigManager.GetISPsConfig()
	if ispsIds != nil && len(ispsIds) != 0 {
		for _, ispId := range ispsIds {
			ispCfg, exist := ispsCfg[ispId]
			if !exist {
				errs = append(errs, fmt.Errorf("unknown isp: %s", ispId))
				continue
			}
			ispsToUpdate = append(ispsToUpdate, &ispCfgWithId{
				ispId:  ispId,
				ispCfg: ispCfg,
			})
		}
	} else {
		for ispId, ispCfg := range ispsCfg {
			ispsToUpdate = append(ispsToUpdate, &ispCfgWithId{
				ispId:  ispId,
				ispCfg: ispCfg,
			})
		}
	}

	errs = append(errs, updatePublicIps(ispsToUpdate)...)

	return errors.Join(errs...)
}

const noData = "<NO-DATA>"

func updatePublicIps(ispsToUpdate []*ispCfgWithId) []error {
	var errs []error
	dao := isppublicip.New()
	for _, ispToUpdate := range ispsToUpdate {
		ipInfo, err := dao.GetByISP(ispToUpdate.ispId)
		if err != nil {
			errs = append(errs, fmt.Errorf("error updating ip for %s. %w", ispToUpdate.ispId, err))
			continue
		}

		var localIp = noData
		if ipInfo != nil {
			localIp = ipInfo.IP
		}

		currentIp, err := GetPublicIP(ispToUpdate.ispCfg)
		if err != nil {
			errs = append(errs, fmt.Errorf("error updating ip for %s. %w", ispToUpdate.ispId, err))
			continue
		}

		environment.Get().OutputLogger.Printf("Provider: %s Locally saved public IP: %s, Current public IP: %s\n",
			ispToUpdate.ispCfg.Name,
			localIp,
			currentIp)

		if ipInfo == nil {
			err = dao.Insert(&entities.ISPPublicIP{
				ISP:      ispToUpdate.ispId,
				IP:       currentIp,
				Modified: time.Now(),
			})
			if err != nil {
				errs = append(errs, fmt.Errorf("error updating ip for %s. %w", ispToUpdate.ispId, err))
				continue
			}
		} else if ipInfo.IP != currentIp {
			ipInfo.IP = currentIp
			ipInfo.Modified = time.Now()
			err = dao.UpdateIpInfo(ipInfo)
			if err != nil {
				errs = append(errs, fmt.Errorf("error updating ip for %s. %w", ispToUpdate.ispId, err))
				continue
			}
		}
	}
	return errs
}
