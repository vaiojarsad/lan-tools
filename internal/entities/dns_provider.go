package entities

func NewDnsProvider(storageId int64, code, name, serviceType string, serviceCfg map[string]string) *DnsProvider {
	return &DnsProvider{
		storageId:   storageId,
		Code:        code,
		Name:        name,
		ServiceType: serviceType,
		ServiceCfg:  serviceCfg,
	}
}

type DnsProvider struct {
	storageId   int64
	Code        string
	Name        string
	ServiceType string
	ServiceCfg  map[string]string
}

func (s *DnsProvider) StorageId() int64 {
	return s.storageId
}
