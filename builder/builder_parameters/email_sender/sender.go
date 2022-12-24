package email_sender

import (
	"patterns/builder/builder_parameters/email_builder"
	"patterns/builder/builder_parameters/entity"
)

func sendMailTemplate(email *entity.Email) error { return nil }

type build func(*email_builder.EmailBuilder)

func SendEmail(build build) error {
	builder := email_builder.EmailBuilder{}
	build(&builder)
	email := builder.Email
	return sendMailTemplate(&email)
}
