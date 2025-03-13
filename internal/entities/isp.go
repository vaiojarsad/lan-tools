package entities

import "time"

func NewIsp(storageId int64, code, name, publicIpGetterType string, publicIpGetterCfg map[string]string, publicIp string,
	publicIpModTime time.Time) *Isp {
	return &Isp{
		storageId:          storageId,
		Code:               code,
		Name:               name,
		PublicIpGetterType: publicIpGetterType,
		PublicIpGetterCfg:  publicIpGetterCfg,
		PublicIp:           publicIp,
		PublicIpModTime:    publicIpModTime,
	}
}

type Isp struct {
	storageId          int64
	uuid               string
	Code               string
	Name               string
	PublicIpGetterType string
	PublicIpGetterCfg  map[string]string
	PublicIp           string
	PublicIpModTime    time.Time
}

func (i *Isp) StorageId() int64 {
	return i.storageId
}

func (i *Isp) UUID() string {
	return i.uuid
}
