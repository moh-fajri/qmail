package main

import (
	"fmt"
	"path"

	"github.com/moh-fajri/qmail"
)

func main() {
	// set credential aws
	mail := qmail.NewEmailAws(&qmail.EmailAws{
		AccessKey: "accees key",
		SecretKey: "seccret key",
		Region:    "region",
	})
	data := struct {
		Name   string
		Mobile string
	}{
		Name:   "example",
		Mobile: "0899890908",
	}

	emailTo := []string{"user@mail.com", "user@gmail.com"}
	emailCc := []string{"usercc@mail.com"}
	filepath := path.Join("example", "template.html")
	err := mail.SendEmail("noreply@gmail.com", emailTo, emailCc, "subject", filepath, data)
	fmt.Println("Error :", err)
}
