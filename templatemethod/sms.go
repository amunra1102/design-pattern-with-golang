package templatemethod

import "fmt"

type SMS struct {}

func (s *SMS) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("SMS: generating random OTP %s\n", randomOTP)
	return randomOTP
}

func (s *SMS) saveOTPCache(otp string)  {
	fmt.Printf("SMS: saving OTP: %s to cache \n", otp)
}

func (s *SMS) getMessage(otp string) string {
	return "SMS OTP for login is" + otp
}

func (s *SMS) sendNotification(message string) error {
	fmt.Printf("SMS: sending sms %s\n", message)
	return nil
}
