package mediator

import "fmt"

type PassengerTrain struct {
	Mediator Mediator
}

func (p *PassengerTrain) RequestArrival() {
	if p.Mediator.CanLand(p) {
		fmt.Println("Passenger Train: Landing")
	} else {
		fmt.Println("Passenger Train: Waiting")
	}
}

func (p *PassengerTrain) Departure() {
	fmt.Println("Passenger Train: Leaving")
	p.Mediator.NotifyFree()
}

func (p *PassengerTrain) PermitArrival() {
	fmt.Println("Passenger Train: Arrival Permitted. Landing")
}
