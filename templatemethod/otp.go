package templatemethod

type OTP struct {
	ObjectOTP IOtp
}

func (o *OTP) GenAndSendOTP(otpLength int) error {
	otp := o.ObjectOTP.genRandomOTP(otpLength)
	o.ObjectOTP.saveOTPCache(otp)
	message := o.ObjectOTP.getMessage(otp)
	err := o.ObjectOTP.sendNotification(message)
	if err != nil {
		return err
	}

	return nil
}