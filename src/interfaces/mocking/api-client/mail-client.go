// This example shows how to use the real implementation of an API or a mocked version
// Interesting is that the mocking logic is implemented by the client and not by the API provider as per guidelines
// https://github.com/golang/go/wiki/CodeReviewComments#interfaces

// run this command to use the real implementation
// go run ./src/interfaces/mocking/api-client

// run this command to use the mock version
// go run ./src/interfaces/mocking/api-client -mock

package main

import (
	"flag"
	"fmt"

	mail "github.com/EnricoPicci/go-class/src/interfaces/mocking/api-provider"
)

// MailSender is an interface that is implemented by the real API mail.MailSystem
type MailSender interface {
	SendMail(mail.Mail) error
}

// MailSystemMock is the mock implementing the MailSender interface
type MailSystemMock struct {
	mailsSent int
}

func (mS *MailSystemMock) SendMail(m mail.Mail) error {
	mS.mailsSent++
	fmt.Printf("I am the mock simulating to send the mail - I have sent %v mails\n", mS.mailsSent)
	return nil
}

// SendMail sends a mail using the MailSender implementation which is passed to it by the client
func SendMail(m mail.Mail, sender MailSender) error {
	return sender.SendMail(m)
}

// main used the real implementation or the mock implementation of the MailSystem depending on the flag passed from the command line command
func main() {
	mockPtr := flag.Bool("mock", false, "use of the mocked MailSystem")
	flag.Parse()

	var mailSender MailSender
	if *mockPtr {
		mailSender = &MailSystemMock{}
	} else {
		mailSender = mail.NewMailSystem()
	}
	aMail := mail.Mail{Header: "Mail header", Body: "The mail body"}
	err := SendMail(aMail, mailSender)
	if err != nil {
		panic(err)
	}
}
