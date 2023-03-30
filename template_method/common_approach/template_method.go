package template_method

type IOtp interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
}

type Otp struct {
	iOtp IOtp
}

func (o *Otp) GenerateAndSendOTP(length int) error {
	otp := o.iOtp.genRandomOTP(length)
	o.iOtp.saveOTPCache(otp)
	message := o.iOtp.getMessage(otp)
	return o.iOtp.sendNotification(message)
}
