package qmail

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

// EmailAws object
type EmailAws struct {
	AccessKey string
	SecretKey string
	Region    string
}

// configuration handle config aws
func (ea *EmailAws) configuration() (*ses.SES, error) {
	creds := credentials.NewStaticCredentials(ea.AccessKey, ea.SecretKey, "")
	_, err := creds.Get()
	if err != nil {
		return nil, err
	}
	cfg := aws.NewConfig().WithRegion(ea.Region).WithCredentials(creds)
	// Create an SES session.
	return ses.New(session.New(), cfg), nil
}

// parseTemplateHtml to parse template html to string
func (ea *EmailAws) parseTemplateHtml(path string, data interface{}) (string, error) {
	//open and parse a template html file
	var tmpl, err = template.ParseFiles(path)
	// do error stuff here, or log error, etc
	if err != nil {
		return "", err
	}
	//merge template ‘t’ with content
	var doc bytes.Buffer
	err = tmpl.Execute(&doc, data)
	if err != nil {
		return "", err
	}
	return doc.String(), nil
}

// SendEmail to send email with aws
func (ea *EmailAws) SendEmail(emailFrom string, emailTo []string, emailCc []string, subject string, path string, data interface{}) error {
	// parse template html to string
	tmpltString, err := ea.parseTemplateHtml(path, data)
	if err != nil {
		return err
	}
	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: aws.StringSlice(emailTo),
			CcAddresses: aws.StringSlice(emailCc),
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(tmpltString),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(emailFrom),
	}
	// configuration aws
	svc, err := ea.configuration()
	if err != nil {
		return err
	}

	// Attempt to send the email.
	res, err := svc.SendEmail(input)
	// Display error messages if they occur.
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
				return err
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
				return err
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
				return err
			default:
				fmt.Println(aerr.Error())
				return err
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			return err
		}
	}
	fmt.Println(res)
	return nil
}

// NewEmailAws create new EmailAws
func NewEmailAws(sendEmail *EmailAws) *EmailAws {
	return sendEmail
}
