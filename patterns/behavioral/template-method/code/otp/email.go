package otp

import "fmt"

type email struct {
	otp
}

func NewEmail() *email {
	return &email{}
}

func (s *email) GenRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("EMAIL: generaing random otp %s\n", randomOTP)
	return randomOTP
}

func (s *email) SaveOTPCache(otp string) {
	fmt.Printf("EMAIL: saving otp: %s to cache\n", otp)
}

func (s *email) GetMessage(otp string) string {
	return "EMAIL OTP for login is " + otp
}

func (s *email) SendNotification(message string) error {
	fmt.Printf("EMAIL: sending email: %s\n", message)
	return nil
}

func (s *email) PublishMetric() {
	fmt.Printf("EMAIL: publishing metrics\n")
}
