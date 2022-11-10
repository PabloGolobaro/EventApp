package helpers

import (
	"bytes"
	"crypto/tls"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/localconf"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/models"
	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

type EmailData struct {
	URL      string
	Username string
	Subject  string
}

// ? Email template parser

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

func SendEmail(user *models.User, data *EmailData) {

	var body bytes.Buffer

	//template, err := ParseTemplateDir("./static/templates")
	//if err != nil {
	//	log.Fatal("Could not parse template", err)
	//}

	file, err := template.ParseFiles("./static/templates/verificationCode.html")
	if err != nil {
		return
	}
	err = file.ExecuteTemplate(&body, "verificationCode.html", &data)
	if err != nil {
		return
	}

	// Sender data.
	from := localconf.Config.EmailFrom
	smtpPass := localconf.Config.SMTPPass
	smtpUser := localconf.Config.SMTPUser
	to := user.Email
	smtpHost := localconf.Config.SMTPHost
	smtpPort := localconf.Config.SMTPPort
	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send Email
	if err := d.DialAndSend(m); err != nil {
		log.Fatal("Could not send email: ", err)
	}

}

//From ApiClient
//newEmail := &api.Email{
//	From:    localconf.Config.EmailFrom,
//	Name:    "PabloGolobar",
//	Subject: data.Subject,
//	To:      user.Email,
//	To_name: "User",
//	Html:    body.String(),
//	Text:    html2text.HTML2Text(body.String()),
//}
//_, err = localconf.Config.API.MessengerApi.SendEmail(context.Background(), "auth", newEmail)
//if err != nil {
//	return
//}

//SendGrid
//from := mail.NewEmail("PabloGolobar", localconf.Config.EmailFrom)
//subject := data.Subject
//to := mail.NewEmail("User", user.Email)
//plainTextContent := html2text.HTML2Text(body.String())
//htmlContent := body.String()
//message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
//client := sendgrid.NewSendClient(localconf.Config.SENDGRID_API_KEY)
//response, err := client.Send(message)
//if err != nil {
//	log.Println(err)
//} else {
//	log.Println(response.StatusCode)
//}
