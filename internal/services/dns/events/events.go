package events

import (
	"github.com/vaiojarsad/lan-tools/internal/environment"
	"github.com/vaiojarsad/lan-tools/internal/events"
	"github.com/vaiojarsad/lan-tools/internal/services/dns"
)

func HandleISPPublicIPCheck(e *events.ISPPublicIPCheck) error {
	environment.Get().OutputLogger.Printf("DNS ISP public IP check event handler\n")
	return dns.SyncRecordsAWithIP(e.ISP, e.NewIP)
}
