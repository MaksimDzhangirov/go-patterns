package main

import (
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/chain-of-responsibility/code/hospital"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/chain-of-responsibility/code/patient"
)

func main() {
	cashier := hospital.NewCashier()
	// Определяем куда отправить пациент после комнаты выдачи медикаментов
	medical := hospital.NewMedical()
	medical.SetNext(cashier)
	// Определяем куда отправить пациент после кабинета врача
	doctor := hospital.NewDoctor()
	doctor.SetNext(medical)
	// Определяем куда отправить пациент после регистратуры
	reception := hospital.NewReception()
	reception.SetNext(doctor)
	hospitalPatient := patient.NewPatient("abc")
	// Пациент приходит в больницу
	reception.Execute(hospitalPatient)
}
