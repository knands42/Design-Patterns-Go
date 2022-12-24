package email_builder

import (
	"patterns/builder/builder_parameters/entity"
	"strings"
)

type EmailBuilder struct {
	Email entity.Email
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
	if strings.Contains(from, "@") == false {
		panic("Invalid email address")
	}

	b.Email.From = from
	return b
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
	if strings.Contains(to, "@") == false {
		panic("Invalid email address")
	}

	b.Email.To = to
	return b
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.Email.Subject = subject
	return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.Email.Body = body
	return b
}
