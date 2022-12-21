package email_builder

import (
	"patterns/builder/email_builder_parameters/email_builder"
	"patterns/builder/email_builder_parameters/email_sender"
	"testing"
)

func Test(t *testing.T) {
	email_sender.SendEmail(func(b *email_builder.EmailBuilder) {
		b.From("foo@bar.com").
			To("bar@baz.com").
			Subject("Meeting").
			Body("Hello, do you want to meet?")
	})
}
