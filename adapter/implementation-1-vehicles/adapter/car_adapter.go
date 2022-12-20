package adapter

import (
	"fmt"
	"patterns/adapter/implementation-1-vehicles/vehicles"
)

type CarAdapter struct {
	CarTransport *vehicles.Car
}

func (c *CarAdapter) NavigateToDestination() {
	fmt.Println("Adapter modify car to allow navigation")
	c.CarTransport.DriveToDestination()
}
