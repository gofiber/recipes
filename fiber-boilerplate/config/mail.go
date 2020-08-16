package config

import (
	"github.com/gofiber/fiber"
	. "github.com/itsursujit/fiber-boilerplate/app"
	"github.com/itsursujit/fiber-boilerplate/mail"
	"github.com/valyala/bytebufferpool"
	"log"
	"time"
)

type MailConfiguration struct {
	Mail_Host         string
	Mail_Port         int
	Mail_Username     string
	Mail_Password     string
	Mail_Encryption   string
	Mail_From_Address string
	Mail_From_Name    string
}

var MailConfig *MailConfiguration //nolint:gochecknoglobals

func LoadMailConfig() {
	loadDefaultMailConfig()
	ViperConfig.Unmarshal(&MailConfig)
}

func loadDefaultMailConfig() {
	ViperConfig.SetDefault("MAIL_HOST", "smtp.mailtrap.io")
	ViperConfig.SetDefault("MAIL_PORT", "2525")
	ViperConfig.SetDefault("MAIL_USERNAME", "2c62400cd5524d")
	ViperConfig.SetDefault("MAIL_PASSWORD", "a6770ce74dfee4")
	ViperConfig.SetDefault("MAIL_ENCRYPTION", "tls")
	ViperConfig.SetDefault("MAIL_FROM_ADDRESS", "itsursujit@gmail.com")
	ViperConfig.SetDefault("MAIL_FROM_NAME", "fiber-boilerplate")
}

func Send(to string, subject string, body string, cc string, from string) {
	if MailerServer == nil {
		SetupMailer()
	}
	//New email simple html with inline and CC
	email := mail.NewMSG()
	if from == "" { //nolint:wsl
		from = "Sujit Baniya <itsursujit@gmail.com>"
	}
	email.SetFrom(from). //nolint:wsl
				AddTo(to).
				SetSubject(subject)
	if cc != "" { //nolint:wsl
		email.AddCc(cc)
	}
	email.SetBody(mail.TextHTML, body) //nolint:wsl

	//Call Send and pass the client
	err := email.Send(Mailer)

	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email Sent")
	}
}

func PrepareHtml(view string, body fiber.Map) string {
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)
	// app.Settings.Views.Render
	if err := TemplateEngine.Render(buf, view, body, "layouts/email"); err != nil {
		// handle err
	}
	return buf.String()
}

func SetupMailer() {
	LoadMailConfig()
	var err error
	MailerServer = mail.NewSMTPClient()
	MailerServer.Host = MailConfig.Mail_Host
	MailerServer.Port = MailConfig.Mail_Port
	MailerServer.Username = MailConfig.Mail_Username
	MailerServer.Password = MailConfig.Mail_Password
	if MailConfig.Mail_Encryption == "tls" {
		MailerServer.Encryption = mail.EncryptionTLS
	} else {
		MailerServer.Encryption = mail.EncryptionSSL
	}

	//Variable to keep alive connection
	MailerServer.KeepAlive = false

	//Timeout for connect to SMTP Server
	MailerServer.ConnectTimeout = 10 * time.Second

	//Timeout for send the data and wait respond
	MailerServer.SendTimeout = 10 * time.Second
	Mailer, err = MailerServer.Connect()
	if err != nil {
		log.Print(err)
	}
}
