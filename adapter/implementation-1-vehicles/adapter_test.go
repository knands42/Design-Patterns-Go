package implementation_1_vehicles

import (
	"patterns/adapter/implementation-1-vehicles/adapter"
	"patterns/adapter/implementation-1-vehicles/client"
	"patterns/adapter/implementation-1-vehicles/vehicles"
	"testing"
)

func Test_(t *testing.T) {
	client := &client.Client{}
	boat := &vehicles.Boat{}

	client.StartNavigation(boat)

	car := &vehicles.Car{}
	carAdapter := &adapter.CarAdapter{
		CarTransport: car,
	}

	client.StartNavigation(carAdapter)

}
