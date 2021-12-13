package main

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/template-method/code/otp"
	"log"
)

func main() {
	smsOTP := otp.NewSms()
	o := otp.NewOTP(smsOTP)
	err := o.GenAndSendOTP(4)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("")
	emailOTP := otp.NewEmail()
	o = otp.NewOTP(emailOTP)
	err = o.GenAndSendOTP(4)
	if err != nil {
		log.Fatalln(err)
	}
}