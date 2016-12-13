package model

import (
	"net/smtp"
	"strings"
)

const (
	HOST        = "smtp.163.com"
	SERVER_ADDR = "smtp.163.com:25"
	ALIAS       = "叁公子的还书提醒助手"
	SUBJECT     = "还书提醒"
	USER        = "ustb_books_alert@163.com" //发送邮件的邮箱
	PASSWORD    = "ustb806"                  //发送邮件邮箱的密码
)

type Email struct {
	Alias   string
	To      string
	Subject string
	Content string
}

func NewEmail(to, content string) *Email {
	alias := ALIAS + "<" + USER + ">"
	return &Email{Alias: alias, To: to, Subject: SUBJECT, Content: content}
}

func SendEmail(email *Email) error {
	auth := smtp.PlainAuth("", USER, PASSWORD, HOST)

	str := strings.Replace("From: "+email.Alias+"~To: "+email.To+"~Subject: "+email.Subject+"~~", "~", "\r\n", -1) + email.Content

	err := smtp.SendMail(
		SERVER_ADDR,
		auth,
		USER,
		[]string{email.To},
		[]byte(str),
	)

	return err
}
