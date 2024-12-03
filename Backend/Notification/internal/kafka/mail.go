package kafka

import (
	"fmt"
	"log/slog"
	"net/smtp"
	"strconv"
)

var (
	CheckTemplate = `
	<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body style="font-family: Arial, sans-serif; margin: 0; padding: 0; background-color: #f4f4f4; text-align: center;">

<div style="display: block; width: 100%; text-align: center; margin-bottom: 80px;">
    <div style="background: #1969B3; padding: 20px; margin-top: 30px; display: inline-block; text-align: center; border-radius: 20px;">
        <p style="color: white; font-size: 32px; margin: 0; font-family: Arial, sans-serif;">
            SkufSchool
        </p>
    </div>
</div>

<div style="background: #01182E; display: block; width:500px; margin: 50px auto; border-radius: 30px; padding: 10px 30px; font-family: Arial, sans-serif;">
    <div style="display: inline-block; padding: 10px 0;">
        <h2 style="font-size: 32px; color: white; margin: 0;">
            Почти готово!
        </h2>
    </div>
    <div style="margin-bottom: 90px;">
        <p style="color: white; font-size: 20px; margin: 0;">
            Для завершения регистрации введите код в приложении:
        </p>
    </div>
    
    <table align="center" style="margin: 0 auto;">
        <tr>
            %s
        </tr>
    </table>
</div>

<div style="display: flex; justify-content: center;">
    <p style="font-size: 20px; font-family: Arial, sans-serif;">
        С уважением, команда <span style="font-weight: bold;">SkufSchool</span>
    </p>
</div>

</body>
</html>
`
	elementTemplate = `<td style="border-radius: 15px; background: #D9D9D9; width: 100px; height: 100px; text-align: center;">
                <p style="font-size: 64px; margin: 0;">%s</p>
            </td>`
)

const (
	smtpHost = "smtp.gmail.com"
	smtpPort = "587"
)

func (r *KafkaReader) SendMail(to string, checkInt int) error {
	checkValue := strconv.Itoa(checkInt)
	if len(checkValue) != 4 {
		slog.Error("incorrect checkInt value")
		return nil
	}

	body := ""

	for i := range checkValue {
		element := fmt.Sprintf(elementTemplate, string(checkValue[i]))
		body += element
	}

	subject := "Код для регистрации в SkufSchool"
	message := fmt.Sprintf(CheckTemplate, "", body)
	auth := smtp.PlainAuth("", r.MailConfig.Mail, r.MailConfig.Password, smtpHost)
	resultMsg := fmt.Sprintf(
		"From: %s\nTo: %s\nSubject: %s\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n%s",
		r.MailConfig.Mail, to, subject, message)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, r.MailConfig.Mail, []string{to}, []byte(resultMsg))
	if err != nil {
		return err
	}
	return nil
}
