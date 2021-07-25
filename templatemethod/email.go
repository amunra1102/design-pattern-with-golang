package templatemethod

import "fmt"

type Email struct {}

func (s *Email) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("Email: generating random OTP %s\n", randomOTP)
	return randomOTP
}

func (s *Email) saveOTPCache(otp string)  {
	fmt.Printf("Email: saving OTP: %s to cache \n", otp)
}

func (s *Email) getMessage(otp string) string {
	return "SMS OTP for login is" + otp
}

func (s *Email) sendNotification(message string) error {
	fmt.Printf("Email: sending sms %s\n", message)
	return nil
}
