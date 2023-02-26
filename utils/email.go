package utils

import (
	"fmt"
	"html/template"
	"net/smtp"
	"bytes"
)
	// Define the email content
	type EmailContent = struct {
		From      string
		To        string
		Subject   string
		Recipient string
		Body      string
		Sender    string
	}



func SendEmail(emailContent EmailContent) error {
	// Define the email template
	templateString := `
		From: {{.From}}
		To: {{.To}}
		Subject: {{.Subject}}

		Dear {{.Recipient}},

		{{.Body}}

		Thamks,
		{{.Sender}}
	`

	// Parse the email template
	tmpl, err := template.New("emailTemplate").Parse(templateString)
	if err != nil {
		fmt.Println("Error parsing email template:", err)
		return err
	}

	// Execute the email template with the email content
	var emailBodyContent bytes.Buffer
	err = tmpl.Execute(&emailBodyContent, emailContent)
	if err != nil {
		fmt.Println("Error executing email template:", err)
		return err
	}

	// Define the email message
	message := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", emailContent.To, emailContent.Subject, emailBodyContent.String()))

	// Define the SMTP server configuration
	auth := smtp.PlainAuth("", "sender@example.com", "password", "smtp.example.com")

	// Send the email
	err = smtp.SendMail("smtp.example.com:587", auth, "sender@example.com", []string{emailContent.To}, message)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return err
	}

	return nil
}