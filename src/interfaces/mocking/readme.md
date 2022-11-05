# Mocking

This is a basic example of how to use interfaces on the api-client to provide a mock implementation of an API provided by a function exposed by the api-provider

## api-provider (mail package)

The api-provider is a package that exposed a type, `MailSystem`, which implements the API method `func (mS *MailSystem) SendMail(m Mail) error`.

There is no concept of mock in the api-provider, just the real implementation

## api-client

The api-client package defines

- an interface, `MailSender`, which defines a method with the same signature as the one exposed by the real api `mail.MailSystem`, i.e. `SendMail(m Mail) error`
- a function, `SendMail(m mail.Mail, sender MailSender) error`, that accepts a value of type `MailSender`, i.e. of the interface defined by the api-client package, and uses this value to execute the actual send operation
- `MailSenderMock`, a mock implementation of `MailSender` interface

Now, with all this machinery implemented, the api-client can decide to

- use the real implementation, passing to the function `SendMail(m mail.Mail, sender MailSender) error` a value of type `mail.MailSystem`
- use the mock implementation, passing to the function `SendMail(m mail.Mail, sender MailSender) error` a value of type `MailSenderMock`

In this way, the api-client can do mocking in a way decoupled from the actual implementation of the api-provider

## Run the example

Run the example against the real implementation (non mock) with the command

`go run ./src/interfaces/mocking/api-client`

Run the example against the mock implementation with the command
`go run ./src/interfaces/mocking/api-client -mock`
