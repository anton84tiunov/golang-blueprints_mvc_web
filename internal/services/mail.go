package services

import (
	"crypto/tls"
	"fmt"
	"net/smtp"

	config "../../internal/config"
)

func SendMessageMail(to_str string, msg string, title string) {
	cfg := config.GLOBAL_CONFIG

	var auth = smtp.PlainAuth("", cfg.SmtpServer.From, cfg.SmtpServer.Password, cfg.SmtpServer.Host)
	var conf = &tls.Config{ServerName: cfg.SmtpServer.Host}
	var conn, err = tls.Dial("tcp", fmt.Sprintf("%s:%d", cfg.SmtpServer.Host, cfg.SmtpServer.Port), conf)

	message := "From: " + cfg.SmtpServer.From + "\n" +
		"To: " + to_str + "\n" +
		"Subject: " + title + " \n\n" +
		msg

	if err != nil {
		fmt.Println("err", err)
		return
	}
	var cl, err1 = smtp.NewClient(conn, cfg.SmtpServer.Host)
	if err1 != nil {
		fmt.Println("err1", err1)
		return
	}
	err2 := cl.Auth(auth)
	if err2 != nil {
		fmt.Println("err2", err2)
		return
	}
	err3 := cl.Mail(cfg.SmtpServer.From)
	if err3 != nil {
		fmt.Println("err3", err3)
		return
	}
	err4 := cl.Rcpt(to_str)
	if err4 != nil {
		fmt.Println("err4", err4)
		return
	}
	var w, err5 = cl.Data()
	if err3 != nil {
		fmt.Println("err5", err5)
		return
	}
	_, err6 := w.Write([]byte(message))
	if err6 != nil {
		fmt.Println("err6", err6)
		return
	}
	err7 := w.Close()
	if err7 != nil {
		fmt.Println("err7", err7)
		return
	}
	err8 := cl.Quit()
	if err8 != nil {
		fmt.Println("err8", err8)
		return
	}
}
