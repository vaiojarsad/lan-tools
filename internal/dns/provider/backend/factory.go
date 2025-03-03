package backend

func NewDNSProviderBackendService(resType string, resCfg map[string]string) (Service, error) {
	if resType == "cloudflare" {
		return newCloudflareDNSService(resCfg)
	}
	return nil, nil
}
