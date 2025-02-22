package isppublicip

import "time"

func New(storageId int64, ISP, IP string, modified time.Time) *ISPPublicIP {
	return &ISPPublicIP{
		storageId: storageId,
		ISP:       ISP,
		IP:        IP,
		Modified:  modified,
	}
}

type ISPPublicIP struct {
	storageId int64
	ISP       string
	IP        string
	Modified  time.Time
}

func (i *ISPPublicIP) StorageId() int64 {
	return i.storageId
}
