# qmail

[![Build Status](https://travis-ci.org/moh-fajri/qmail.svg?branch=master)](https://travis-ci.com/moh-fajri/qmail)

qmail is send email using aws ses

## Instalation

When used with Go modules, use the following import path:

```
go get -u github.com/moh-fajri/qmail
```

## Basic Usage

send email using aws ses
```go
import "github.com/moh-fajri/qmail"

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
// send email 
err := mail.SendEmail("noreply@gmail.com", emailTo, emailCc, "subject", filepath, data)
```
