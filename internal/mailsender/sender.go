package mailsender

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/smtp"

	"github.com/rs/zerolog/log"
)

type Sender struct {
	senderEmail string
	password    string
	smtpHost    string
	smtpPort    string
}

type SenderDeps struct {
	SenderEmail string
	Password    string
	SMTPHost    string
	SMTPPort    string
}

func NewSender(deps *SenderDeps) *Sender {
	return &Sender{
		senderEmail: deps.SenderEmail,
		password:    deps.Password,
		smtpHost:    deps.SMTPHost,
		smtpPort:    deps.SMTPPort,
	}
}

func (s *Sender) SendEmail(ctx context.Context, receiver, payload, subject string) error {
	to := []string{
		receiver,
	}

	header := make(map[string]string)
	header["From"] = s.senderEmail
	//header["To"] = to.String()
	header["Subject"] = subject
	//header["MIME-Version"] = "1.0"
	//header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(payload))

	auth := smtp.PlainAuth("", s.senderEmail, s.password, s.smtpHost)

	err := smtp.SendMail(s.smtpHost+":"+s.smtpPort, auth, s.senderEmail, to, []byte(message))
	if err != nil {
		fmt.Println(err)
		return err
	}
	log.Info().Msgf("email sent to %s", receiver)
	return nil
}
