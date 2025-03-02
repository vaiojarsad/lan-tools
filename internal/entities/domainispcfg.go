package entities

func NewDomainISPCfg(domainId, ispId int64, dnsProviderCurrentIp, dnsProviderRecordId string) *DomainISPCfg {
	return &DomainISPCfg{
		DomainId:             domainId,
		ISPId:                ispId,
		DnsProviderCurrentIp: dnsProviderCurrentIp,
		DnsProviderRecordId:  dnsProviderRecordId,
	}
}

type DomainISPCfg struct {
	DomainId             int64
	ISPId                int64
	DnsProviderCurrentIp string
	DnsProviderRecordId  string
}
