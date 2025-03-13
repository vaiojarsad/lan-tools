package entities

func NewDnsState(domainId, ispId int64, dnsProviderCurrentIp, dnsProviderRecordId, dnsProviderSyncStatus string) *DnsState {
	return &DnsState{
		DomainId:              domainId,
		IspId:                 ispId,
		DnsProviderCurrentIp:  dnsProviderCurrentIp,
		DnsProviderRecordId:   dnsProviderRecordId,
		DnsProviderSyncStatus: dnsProviderSyncStatus,
	}
}

type DnsState struct {
	DomainId              int64
	IspId                 int64
	DnsProviderCurrentIp  string
	DnsProviderRecordId   string
	DnsProviderSyncStatus string
}
