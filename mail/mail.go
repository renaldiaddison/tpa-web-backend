package mail

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

const SMTP_HOST = "smtp.gmail.com"
const SMTP_PORT = 587
const SENDER_NAME = "linkHEdIn <linkhedin@gmail.com>"
const SENDER_EMAIL = "renacierr@gmail.com"
const SENDER_EMAIL_PASSWORD = "txqkhfsshskjkfyk"

func SendEmail(text string, subject string, to string, link string) {

	mail := gomail.NewMessage()

	mail.SetHeader("From", SENDER_NAME)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", text+link)

	dial := gomail.NewDialer(SMTP_HOST, SMTP_PORT, SENDER_EMAIL, SENDER_EMAIL_PASSWORD)

	if err := dial.DialAndSend(mail); err != nil {
		fmt.Println(err)
		panic(err)
	}

}

// func SendEmailForgotPassword(userEmail string, link string) {

// 	mail := gomail.NewMessage()

// 	mail.SetHeader("From", SENDER_NAME)
// 	mail.SetHeader("To", userEmail)
// 	mail.SetHeader("Subject", "linkHEdIn Forgot Password")
// 	mail.SetBody("text/html", "This is your password reset link!! "+link)

// 	dial := gomail.NewDialer(SMTP_HOST, SMTP_PORT, SENDER_EMAIL, SENDER_EMAIL_PASSWORD)

// 	if err := dial.DialAndSend(mail); err != nil {
// 		fmt.Println(err)
// 		panic(err)
// 	}

// }
