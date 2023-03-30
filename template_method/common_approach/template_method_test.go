package template_method

import "testing"

func Test_TemplateMethod(t *testing.T) {
	smsOtp := &Sms{}
	otp := Otp{
		iOtp: smsOtp,
	}
	otp.GenerateAndSendOTP(4)

	emailOtp := &Email{}
	otp = Otp{
		iOtp: emailOtp,
	}
	otp.GenerateAndSendOTP(4)
}
