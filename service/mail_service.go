package service

import (
	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendEmail(to string, subject string, message string) error {

	m := gomail.NewMessage()

	m.SetHeader("From", os.Getenv("SMTP_USER"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("Text/html", message)

	from := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_PASSWORD")
	host := os.Getenv("SMTP_HOST")
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

	d := gomail.NewDialer(host, port, from, password)

	err := d.DialAndSend(m)
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	return err

	// msg := "Message to you!!!"
	// body := []byte(msg)

	// auth := smtp.PlainAuth("", user, password, host)

	// err := smtp.SendMail(host+":"+port, auth, from, toList, body)
	// if err != nil {
	// 	log.Printf("%v", err)
	// 	return err
	// }

	// log.Printf("Successfully sent email to %v \n", toList)

	// return err
}
