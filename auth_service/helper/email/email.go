package email

import (
	"authservice/app/config"
	"authservice/internal"
	"bytes"
	"crypto/tls"
	"text/template"

	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
)

type emailData struct {
	URL     string
	Name    string
	Subject string
}

type emailService struct {
	url      string
	from     string
	password string
	host     string
	port     int
	user     string
}

type EmailInterface interface {
	SendResetPasswordLink(user *internal.User, token string) error
	SendVerificationLink(user *internal.User, token string) error
	SendCodeResetPassword(user *internal.User, code string) error
	SendCodeResetEmail(user *internal.User, code string) error
}

func New() EmailInterface {
	cfg := config.InitConfig()
	emailCfg := cfg
	return &emailService{
		url:      emailCfg.PASSWD_URL,
		from:     emailCfg.EMAIL_FROM,
		user:     emailCfg.SMTP_USER,
		password: emailCfg.SMTP_PASS,
		host:     emailCfg.SMTP_HOST,
		port:     emailCfg.SMTP_PORT,
	}
}

func (e *emailService) SendResetPasswordLink(user *internal.User, token string) error {
	t, err := template.ParseGlob("utils/templates/resetpasswordlink.html")
	if err != nil {
		return err
	}

	to := user.Email
	data := &emailData{
		URL:     e.url + "/reset-password?token=" + token,
		Name:    user.FullName,
		Subject: "Reset Password",
	}

	var body bytes.Buffer
	if err := t.Execute(&body, data); err != nil {
		return err
	}

	m := gomail.NewMessage()

	m.SetHeader("From", e.from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(e.host, e.port, e.user, e.password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send Email
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func (e *emailService) SendVerificationLink(user *internal.User, token string) error {
	t, err := template.ParseGlob("utils/templates/verifiedlink.html")
	if err != nil {
		return err
	}

	to := user.Email
	data := &emailData{
		URL:     e.url + "/verification?token=" + token,
		Name:    user.FullName,
		Subject: "Email Verification",
	}

	var body bytes.Buffer
	if err := t.Execute(&body, data); err != nil {
		return err
	}

	m := gomail.NewMessage()

	m.SetHeader("From", e.from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(e.host, e.port, e.user, e.password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send Email
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func (e *emailService) SendCodeResetPassword(user *internal.User, code string) error {
	t, err := template.ParseGlob("utils/templates/resetpasswordcode.html")
	if err != nil {
		return err
	}

	data := &emailData{
		URL:     code,
		Name:    user.FullName,
		Subject: "Reset Password Code",
	}

	var body bytes.Buffer
	if err := t.Execute(&body, data); err != nil {
		return err
	}

	m := gomail.NewMessage()

	m.SetHeader("From", e.from)
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(e.host, e.port, e.user, e.password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send Email
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

// SendCodeResetEmail implements EmailInterface.
func (e *emailService) SendCodeResetEmail(user *internal.User, code string) error {
	t, err := template.ParseGlob("utils/templates/resetpasswordcode.html")
	if err != nil {
		return err
	}

	data := &emailData{
		URL:     code,
		Name:    user.FullName,
		Subject: "Verified Email Code",
	}

	var body bytes.Buffer
	if err := t.Execute(&body, data); err != nil {
		return err
	}

	m := gomail.NewMessage()

	m.SetHeader("From", e.from)
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(e.host, e.port, e.user, e.password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send Email
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
