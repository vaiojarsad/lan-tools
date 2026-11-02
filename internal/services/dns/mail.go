package dns

import (
	"bytes"
	"embed"
	"text/template"

	"github.com/go-mail/mail"

	"github.com/vaiojarsad/lan-tools/internal/environment"
)

//go:embed templates/*.tmpl
var templatesFS embed.FS

var mailTemplates = template.Must(template.New("").
	Funcs(template.FuncMap{"inc": func(i int) int { return i + 1 }}).
	ParseFS(templatesFS, "templates/*.tmpl"))

type ipChangedMailData struct {
	Provider  string
	LastIP    string
	CurrentIP string
}

type syncErrorsMailData struct {
	Provider  string
	CurrentIP string
	Errors    []error
}

func sendIpChangedMail(provider, currentIP, lastIP string) error {
	return sendMail("Public IP for "+provider, "ip_changed.tmpl", ipChangedMailData{
		Provider:  provider,
		LastIP:    lastIP,
		CurrentIP: currentIP,
	})
}

func sendSyncErrorsMail(provider, currentIP string, errs []error) error {
	return sendMail("Errors syncing A records for "+provider, "sync_errors.tmpl", syncErrorsMailData{
		Provider:  provider,
		CurrentIP: currentIP,
		Errors:    errs,
	})
}

func sendMail(subject, templateName string, data any) error {
	var body bytes.Buffer
	if err := mailTemplates.ExecuteTemplate(&body, templateName, data); err != nil {
		return err
	}

	c := environment.Get().ConfigManager.GetSMTPConfig()
	m := mail.NewMessage()
	m.SetHeader("From", c.Sender)
	m.SetHeader("To", c.To)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body.String())
	d := mail.NewDialer(c.Host, c.Port, c.Sender, c.Pass)
	return d.DialAndSend(m)
}
