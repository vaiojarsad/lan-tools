package entities

func NewDomain(storageId int64, name, description string, dnsProvider *DNSProvider) *Domain {
	return &Domain{
		storageId:   storageId,
		Name:        name,
		Description: description,
		DnsProvider: dnsProvider,
	}
}

type Domain struct {
	storageId   int64
	Name        string
	Description string
	DnsProvider *DNSProvider
}

func (s *Domain) StorageId() int64 {
	return s.storageId
}

func (s *Domain) DnsProviderStorageId() int64 {
	return s.DnsProvider.StorageId()
}
