package backend

import (
	"context"
	"github.com/vaiojarsad/lan-tools/internal/services/dns/provider/backend/entities"
	"sync"
	"time"

	"github.com/cloudflare/cloudflare-go"

	"github.com/vaiojarsad/lan-tools/internal/utils"
)

func newCloudflareDNSService(cfg map[string]string) (*cloudFlareDNSService, error) {
	token := cfg["token"]
	api, err := cloudflare.NewWithAPIToken(token)
	if err != nil {
		return nil, err
	}

	return &cloudFlareDNSService{
		cfg:       cfg,
		token:     token,
		api:       api,
		zoneCache: make(map[string]string),
	}, nil
}

type cloudFlareDNSService struct {
	cfg          map[string]string
	token        string
	api          *cloudflare.API
	zoneCache    map[string]string
	zoneCacheMtx sync.RWMutex
}

func (s *cloudFlareDNSService) GetRecordsByTypeAndName(zone, rType, name string) ([]*entities.DNSRecord, error) {
	zoneId, err := s.getZoneID(zone)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	records, _, err := s.api.ListDNSRecords(ctx, cloudflare.ZoneIdentifier(zoneId), cloudflare.ListDNSRecordsParams{
		Type: rType,
		Name: name,
	})
	if err != nil {
		return nil, err
	}

	return utils.TransformSlice(records, transformDnsRecord), nil
}

func transformDnsRecord(r cloudflare.DNSRecord) *entities.DNSRecord {
	return &entities.DNSRecord{
		ProviderId: r.ID,
		Type:       r.Type,
		Name:       r.Name,
		Content:    r.Content,
		IspCode:    r.Comment,
	}
}

func (s *cloudFlareDNSService) getZoneID(zoneName string) (string, error) {
	s.zoneCacheMtx.RLock()
	if zoneID, found := s.zoneCache[zoneName]; found {
		s.zoneCacheMtx.RUnlock()
		return zoneID, nil
	}
	s.zoneCacheMtx.RUnlock()

	zoneID, err := s.api.ZoneIDByName(zoneName)
	if err != nil {
		return "", err
	}

	s.zoneCacheMtx.Lock()
	s.zoneCache[zoneName] = zoneID
	s.zoneCacheMtx.Unlock()

	return zoneID, nil
}

func (s *cloudFlareDNSService) UpdateDnsRecord(zone string, record *entities.DNSRecord) error {
	zoneId, err := s.getZoneID(zone)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = s.api.UpdateDNSRecord(ctx, cloudflare.ZoneIdentifier(zoneId), cloudflare.UpdateDNSRecordParams{
		Type:    record.Type,
		Name:    record.Name,
		Content: record.Content,
		ID:      record.ProviderId,
		Comment: &record.IspCode,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *cloudFlareDNSService) CreateDnsRecord(zone string, record *entities.DNSRecord) error {
	zoneId, err := s.getZoneID(zone)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	proxied := false
	r, err := s.api.CreateDNSRecord(ctx, cloudflare.ZoneIdentifier(zoneId), cloudflare.CreateDNSRecordParams{
		CreatedOn:  time.Time{},
		ModifiedOn: time.Time{},
		Type:       record.Type,
		Name:       record.Name,
		Content:    record.Content,
		Comment:    record.IspCode,
		Proxied:    &proxied,
	})
	if err != nil {
		return err
	}
	record.ProviderId = r.ID
	return nil
}
