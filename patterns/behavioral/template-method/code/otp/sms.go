package otp

import "fmt"

type sms struct {
	otp
}

func NewSms() *sms {
	return &sms{}
}

func (s *sms) GenRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("SMS: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *sms) SaveOTPCache(otp string) {
	fmt.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (s *sms) GetMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *sms) SendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}

func (s *sms) PublishMetric() {
	fmt.Printf("SMS: publishing metric\n")
}
