package train

import (
	"fmt"
	m "patterns/mediator"
)

type FreightTrain struct {
	Mediator m.Mediator
}

func (g *FreightTrain) Arrive() {
	if !g.Mediator.CanArrive(g) {
		fmt.Println("FreightTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("FreightTrain: Arrived")
}

func (g *FreightTrain) Depart() {
	fmt.Println("FreightTrain: Leaving")
	g.Mediator.NotifyAboutDeparture()
}

func (g *FreightTrain) PermitArrival() {
	fmt.Println("FreightTrain: Arrival permitted")
	g.Arrive()
}
