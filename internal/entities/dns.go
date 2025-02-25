package entities

func NewDNSProvider(storageId int64, code, name, serviceType string, serviceCfg map[string]string) *DNSProvider {
	return &DNSProvider{
		storageId:   storageId,
		Code:        code,
		Name:        name,
		ServiceType: serviceType,
		ServiceCfg:  serviceCfg,
	}
}

type DNSProvider struct {
	storageId   int64
	Code        string
	Name        string
	ServiceType string
	ServiceCfg  map[string]string
}

func (s *DNSProvider) StorageId() int64 {
	return s.storageId
}
