package utils

import (
	"crypto/tls"
	"fmt"
	"os"

	gomail "gopkg.in/mail.v2"
)

func SendTokenEmail(email, token string) error {
	// Set up the email
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_USER"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Your confimation link")
	body := fmt.Sprintf(`
        <h2>Hello,</h2>
        <p>Click the link below to confirm your email:</p>
        <a href="http://localhost:5000/users/verify/%v">confirm email</a>
        <br><br>
        <p>Thank you!</p>
    `, token)

	// Set the body content as HTML
	m.SetBody("text/html", body)

	// Set up the SMTP server
	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), 587, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASS"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func SendForgotPasswordTokenEmail(email, token string) error {
	// Set up the email
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_USER"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Your confimation link")
	body := fmt.Sprintf(`
        <h2>Hello,</h2>
        <p>Click the link below to reset your password:</p>
        <a href="http://localhost:5000/users/password-reset/%v">confirm email</a>
        <br><br>
        <p>Thank you!</p>
    `, token)

	// Set the body content as HTML
	m.SetBody("text/html", body)

	// Set up the SMTP server
	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), 587, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASS"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
