package entities

import (
	"fmt"
)

type DNSRecord struct {
	ProviderId string
	Type       string
	Name       string
	Content    string
	IspCode    string
}

func (r *DNSRecord) String() string {
	return fmt.Sprintf("{ProviderId: %s, Type: %s, Name: %s, Content: %s, ISP: %s}", r.ProviderId, r.Type, r.Name, r.Content, r.IspCode)
}
