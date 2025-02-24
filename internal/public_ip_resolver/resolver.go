package public_ip_resolver

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/vaiojarsad/lan-tools/internal/utils/httputils"
)

type PublicIPResolver interface {
	Resolve() (string, error)
}

func NewPublicIPResolver(resType string, resCfg map[string]string) (PublicIPResolver, error) {
	if resType == "ipify" {
		u, err := url.Parse(resCfg["url"])
		if err != nil {
			return nil, fmt.Errorf("error parsing URL: %w", err)
		}

		return &ipifyResolver{
			url: u,
			ip:  resCfg["ip"],
		}, nil
	}
	return nil, nil
}

type ipifyResolver struct {
	url *url.URL
	ip  string
}

func (r *ipifyResolver) Resolve() (string, error) {
	u := *r.url
	host := u.Host
	u.Host = r.ip
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), http.NoBody)
	if err != nil {
		return "", fmt.Errorf("error creating http request: %w", err)
	}
	req.Host = host
	client := httputils.CreateCustomHTTPClient(true, host)

	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending http request: %w", err)
	}
	defer func(c io.ReadCloser) {
		_ = c.Close()
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("non expected HTTP status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	return string(data[:]), nil
}

/*
func Resolve() (string, error) {
	fileNameLast, fileNameHist, err := getFileName(provider)
	if err != nil {
		return fmt.Errorf("error getting the file name to store the IP: %w", err)
	}
	lastIP, err := getLastSavedIP(fileNameLast)
	if err != nil {
		return fmt.Errorf("error getting last saved ip: %w", err)
	}
	currentIP, err := getPublicIPFromProvider(provider)
	if err != nil {
		return fmt.Errorf("error getting last saved ip: %w", err)
	}
	if lastIP != currentIP {
		err = saveIP(fileNameLast, currentIP)
		if err != nil {
			return fmt.Errorf("error saving new ip: %w", err)
		}
	}
	if lastIP == "" {
		lastIP = "unknown"
	}

	err = sendMail(provider, lastIP, currentIP)
	if err != nil {
		return fmt.Errorf("error sending mail: %w", err)
	}

	err = saveHist(fileNameHist, currentIP)
	if err != nil {
		return fmt.Errorf("error adding ip to historical data: %w", err)
	}

	return nil
}

func saveIP(fileNameLast, ip string) error {
	return os.WriteFile(fileNameLast, []byte(ip), 00600)
}

func saveHist(fileNameHist, ip string) error {
	f, err := os.OpenFile(fileNameHist, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("%s,%s\n", ip, time.Now().Format(time.RFC3339)))
	if err != nil {
		return err
	}
	return nil
}

func getLastSavedIP(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", err
	}
	return string(data[:]), nil
}

func getPublicIPFromProvider(provider string) (string, error) {
		e := environment.Create()
		url := e.ConfigManager.GetResolverConfig().URL
		host := url.Host
		ip, ok := e.ConfigManager.GetResolverConfig().ProviderToIP[provider]
		if !ok {
			return "", fmt.Errorf("configured destination ip wasn't found for provider %s", provider)
		}
		url.Host = ip
		ctx := context.Background()
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), http.NoBody)
		if err != nil {
			return "", fmt.Errorf("error creating http request: %w", err)
		}
		req.Host = host
		client := httputils.CreateCustomHTTPClient(true, host)

		res, err := client.Do(req)
		if err != nil {
			return "", fmt.Errorf("error sending http request: %w", err)
		}
		defer func(c io.ReadCloser) {
			err = c.Close()
			if err != nil {
				le := loggerutils.GetStdErrorLogger()
				le.Fatalf("Error closing response body: %v", err)
			}
		}(res.Body)

		if res.StatusCode != http.StatusOK {
			return "", fmt.Errorf("non expected HTTP status code: %v", res.StatusCode)
		}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return "", fmt.Errorf("error reading response body: %w", err)
		}

		return string(data[:]), nil
	return "", nil
}

func getFileName(provider string) (last, hist string, err error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", "", err
	}
	return filepath.Join(dir, ".public_ip_resolver-"+provider),
		filepath.Join(dir, ".public_ip_resolver-"+provider+"-hist"), nil
}

func sendMail(provider, lastIP, currentIP string) error {
	c := environment.Create().ConfigManager.GetSMTPConfig()
	m := mail.NewMessage()
	m.SetHeader("From", c.Sender)
	m.SetHeader("To", c.To)
	m.SetHeader("Subject", "Public IP for "+provider)
	m.SetBody("text/plain", "Previous IP: "+lastIP+" Current IP: "+currentIP)
	d := mail.NewDialer("smtp.gmail.com", 587, c.Sender, c.Pass)
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
*/
