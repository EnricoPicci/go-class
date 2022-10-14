package mail

import "fmt"

type MailSystem struct{}

func NewMailSystem() *MailSystem {
	return &MailSystem{}
}

func (mS *MailSystem) SendMail(m Mail) error {
	fmt.Printf("The real MailSystem is sending the mail with header '%v' and body '%v'\n", m.Header, m.Body)
	return nil
}

type Mail struct {
	Header string
	Body   string
}
