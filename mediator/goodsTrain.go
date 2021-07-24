package mediator

import "fmt"

type GoodsTrain struct {
	Mediator Mediator
}

func (g *GoodsTrain) RequestArrival() {
	if g.Mediator.CanLand(g) {
		fmt.Println("Good Train: Landing")
	} else {
		fmt.Println("Good Train: waiting")
	}
}

func (g *GoodsTrain) Departure()  {
	fmt.Println("Good Train: Leaving")
	g.Mediator.NotifyFree()
}

func (g *GoodsTrain) PermitArrival() {
	fmt.Println("Good Train: Arrival Permitted. Landing")
}