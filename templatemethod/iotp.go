package templatemethod

type IOtp interface {
	genRandomOTP(len int)string
	saveOTPCache(otp string)
	getMessage(otp string) string
	sendNotification(message string) error
}
