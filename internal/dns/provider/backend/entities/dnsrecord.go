package entities

import "fmt"

type DNSRecord struct {
	ProviderId string
	Type       string
	Name       string
	Content    string
}

func (r *DNSRecord) String() string {
	return fmt.Sprintf("{ProviderId: %s, Type: %s, Name: %s, Content: %s}", r.ProviderId, r.Type, r.Name, r.Content)
}
