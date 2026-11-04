package entities

func NewDNSISP(domainId, ispId int64, dnsProviderRecordId string) *DNSISP {
	return &DNSISP{
		DomainId:            domainId,
		ISPId:               ispId,
		DNSProviderRecordId: dnsProviderRecordId,
	}
}

type DNSISP struct {
	DomainId            int64
	ISPId               int64
	DNSProviderRecordId string
}
