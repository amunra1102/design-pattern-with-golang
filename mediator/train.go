package mediator

type Train interface {
	RequestArrival()
	Departure()
	PermitArrival()
}
