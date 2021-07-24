package chainofresponsibility

type Patient struct {
	Name               string
	isRegistered       bool
	isDoctorChecked    bool
	isMedicineProvided bool
	isPaid             bool
}
