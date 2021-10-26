package main

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"text/template"

	"github.com/joho/godotenv"
)

type data struct {
	email  string
	name   string
	class  string
	mentor string
}

type message struct {
	name string
}

type configuration struct {
	email    string
	password string
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {

	// Sender data.
	from := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")

	// Receiver email address.
	to := []data{
		// {name: "Afandy Wibowo", email: "afandywibowo2000@gmail.com"},
	}

	// smtp server configuration.
	smtpHost := os.Getenv("SMTPHOST")
	smtpPort := os.Getenv("SMTPPORT")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, _ := template.ParseFiles("template.html")

	for _, each := range to {
		var body bytes.Buffer

		mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
		body.Write([]byte(fmt.Sprintf("Subject: Congratulations! \n%s\n\n", mimeHeaders)))

		t.Execute(&body, struct {
			Name string
		}{
			Name: each.name,
		})

		// Sending email.
		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{each.email}, body.Bytes())
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(`Email sent to email:`, each.email, `name:`, each.name)
	}

}
