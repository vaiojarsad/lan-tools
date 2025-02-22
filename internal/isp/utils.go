package isp

import (
	"fmt"
	"time"

	"github.com/vaiojarsad/cloudflare-tools/internal/dao/isppublicip"
	"github.com/vaiojarsad/cloudflare-tools/internal/environment"
)

type Data struct {
	id                   string
	name                 string
	publicIp             string
	publicIpLastModified time.Time
}

func (d *Data) ID() string {
	return d.id
}

func (d *Data) Name() string {
	return d.name
}

func (d *Data) PublicIp() string {
	return d.publicIp
}

func (d *Data) PublicIpLastModified() time.Time {
	return d.publicIpLastModified
}

func (d *Data) String() string {
	if d.publicIp == "" {
		return fmt.Sprintf("Id: %s Name: %s Public IP: Unknown", d.id, d.name)
	}
	return fmt.Sprintf("Id: %s Name: %s Public IP: %s (IP from local cache (last modification %s). Current one "+
		"may differ.)", d.id, d.name, d.publicIp, d.publicIpLastModified.Format(time.RFC3339))
}

func List() ([]*Data, error) {
	ispsCfg := environment.Get().ConfigManager.GetISPsConfig()
	dao := isppublicip.New()
	var ispsData []*Data
	for ispId, ispCfg := range ispsCfg {
		ispData := &Data{
			id:                   ispId,
			name:                 ispCfg.Name,
			publicIp:             "",
			publicIpLastModified: time.Time{},
		}

		ipInfo, err := dao.GetByISP(ispId)
		if err != nil {
			return nil, err
		}
		if ipInfo != nil {
			ispData.publicIp = ipInfo.IP
			ispData.publicIpLastModified = ipInfo.Modified
		}
		ispsData = append(ispsData, ispData)
	}
	return ispsData, nil
}
