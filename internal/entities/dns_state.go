package entities

func NewDnsState(domainId, ispId int64, dnsProviderCurrentIp, dnsProviderRecordId, dnsProviderSyncStatus string) *DnsState {
	return &DnsState{
		DomainId:              domainId,
		ISPId:                 ispId,
		DnsProviderCurrentIp:  dnsProviderCurrentIp,
		DnsProviderRecordId:   dnsProviderRecordId,
		DnsProviderSyncStatus: dnsProviderSyncStatus,
	}
}

type DnsState struct {
	DomainId              int64
	ISPId                 int64
	DnsProviderCurrentIp  string
	DnsProviderRecordId   string
	DnsProviderSyncStatus string
}
