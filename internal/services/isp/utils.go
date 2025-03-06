package isp

import (
	"fmt"
	"github.com/vaiojarsad/lan-tools/internal/utils/public_ip_resolver"
)

func GetPublicIP(getterType string, getterCfg map[string]string) (string, error) {
	r, err := public_ip_resolver.NewPublicIPResolver(getterType, getterCfg)
	if err != nil {
		return "", err
	}
	if r == nil {
		return "", fmt.Errorf("no resolver found for type %s", getterType)
	}
	ip, err := r.Resolve()
	if err != nil {
		return "", err
	}
	return ip, nil
}
