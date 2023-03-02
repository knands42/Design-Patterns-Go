package train_manager

import (
	"patterns/mediator/train"
	"testing"
)

func Test_CanArrive(t *testing.T) {
	stationManager := NewStationManger()

	passengerTrain := train.PassengerTrain{stationManager}
	cargoTrain := train.FreightTrain{stationManager}

	passengerTrain.Arrive()
	cargoTrain.Arrive()

	passengerTrain.Depart()
	cargoTrain.Depart()
}
