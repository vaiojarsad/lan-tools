package entities

import "time"

func NewISP(storageId int64, code, name, publicIpGetterType string, publicIpGetterCfg map[string]string, publicIp string,
	publicIpModTime time.Time) *ISP {
	return &ISP{
		storageId:          storageId,
		Code:               code,
		Name:               name,
		PublicIpGetterType: publicIpGetterType,
		PublicIpGetterCfg:  publicIpGetterCfg,
		PublicIp:           publicIp,
		PublicIpModTime:    publicIpModTime,
	}
}

type ISP struct {
	storageId          int64
	Code               string
	Name               string
	PublicIpGetterType string
	PublicIpGetterCfg  map[string]string
	PublicIp           string
	PublicIpModTime    time.Time
}

func (i *ISP) StorageId() int64 {
	return i.storageId
}
