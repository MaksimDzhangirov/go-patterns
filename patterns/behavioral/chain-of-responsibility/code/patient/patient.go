package patient

type Patient struct {
	name              string
	RegistrationDone  bool
	DoctorCheckUpDone bool
	MedicineDone      bool
	PaymentDone       bool
}

func NewPatient(name string) *Patient {
	return &Patient{name: name}
}