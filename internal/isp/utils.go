package isp

import (
	"fmt"
	"time"

	"github.com/vaiojarsad/lan-tools/internal/public_ip_resolver"
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

type Data struct {
	code                 string
	name                 string
	publicIp             string
	publicIpLastModified time.Time
}

func (d *Data) ID() string {
	return d.code
}

func (d *Data) Name() string {
	return d.name
}

func (d *Data) PublicIp() string {
	return d.publicIp
}

func (d *Data) PublicIpLastModified() time.Time {
	return d.publicIpLastModified
}

func (d *Data) String() string {
	if d.publicIp == "" {
		return fmt.Sprintf("Code: %s Name: %s Public IP: Unknown", d.code, d.name)
	}
	return fmt.Sprintf("Code: %s Name: %s Public IP: %s (IP from local cache (last modification %s). Current one "+
		"may differ.)", d.code, d.name, d.publicIp, d.publicIpLastModified.Format(time.RFC3339))
}
