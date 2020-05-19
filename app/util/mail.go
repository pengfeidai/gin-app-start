package util

import (
	"gin-app-start/app/common"
	"gin-app-start/app/config"
	"strings"

	"gopkg.in/gomail.v2"
)

var logger = common.Logger

// send mail
func SendMail(to, subject, body string) error {
	config := config.Conf
	m := gomail.NewMessage()
	// 设置发件人
	m.SetHeader("From", config.Mail.From)

	// 设置发送给多个用户
	users := strings.Split(to, "")
	m.SetHeader("To", users...)

	// 设置邮件主题
	m.SetHeader("subject", subject)

	// 设置正文
	m.SetBody("text/html", body)
	d := gomail.NewDialer(config.Mail.Host, config.Mail.Port, config.Mail.From, config.Mail.Password)
	if err := d.DialAndSend(m); err != nil {
		logger.Error("SendMail error:", err)
		return err
	}
	return nil
}
