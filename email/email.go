package email

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

const SMTP_HOST = "smtp.gmail.com"
const SMTP_PORT = 587
const SENDER_NAME = "linkHEdIn <linkhedin@gmail.com>"
const SENDER_EMAIL = "addisonrenaldi@gmail.com"
const SENDER_EMAIL_PASSWORD = "jfytkhnqrdsfeasf"

func SendEmail(userEmail string, userId string) {
	mail := gomail.NewMessage()

	mail.SetHeader("From", SENDER_NAME)
	mail.SetHeader("To", userEmail)
	mail.SetHeader("Subject", "linkHEdIn Account Verification")
	link := "http://localhost:5173/" + userId
	mail.SetBody("text/html", "This is your linkHEdIn verification link!! "+link)

	dial := gomail.NewDialer(SMTP_HOST, SMTP_PORT, SENDER_EMAIL, SENDER_EMAIL_PASSWORD)

	if err := dial.DialAndSend(mail); err != nil {
		fmt.Println(err)
		panic(err)
	}

}
