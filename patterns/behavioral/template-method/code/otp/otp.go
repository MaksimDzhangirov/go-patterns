package otp

import "github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/template-method/code/interfaces"

type otp struct {
	iOtp interfaces.IOtp
}

func NewOTP(iOtp interfaces.IOtp) *otp {
	return &otp{
		iOtp: iOtp,
	}
}

func (o *otp) GenAndSendOTP(otpLength int) error {
	otp := o.iOtp.GenRandomOTP(otpLength)
	o.iOtp.SaveOTPCache(otp)
	message := o.iOtp.GetMessage(otp)
	err := o.iOtp.SendNotification(message)
	if err != nil {
		return err
	}
	o.iOtp.PublishMetric()
	return nil
}
