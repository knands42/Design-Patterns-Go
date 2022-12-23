package observer_observable

import (
	"patterns/observer/observer_observable/observable"
	"patterns/observer/observer_observable/observers"
	"testing"
)

func Test_DoctorService(t *testing.T) {
	p := observable.NewPerson("John")
	ds := &observers.DoctorService{}
	p.Subscribe(ds)

	p.CatchACold()
}

func Test_TrafficManagerPropertyObserver(t *testing.T) {
	p := observable.NewPerson("John")
	tm := observers.NewTrafficManagement(p.Observable)
	p.Subscribe(tm)

	p.SetAge(15)
	p.SetAge(16)
	p.SetAge(17)
	p.SetAge(18)
}

func Test_ElectoralRoll(t *testing.T) {
	p := observable.NewPerson("John")
	er := observers.NewElectoralRoll(p.Observable)
	p.Subscribe(er)

	p.SetAge(17)
	p.SetAge(18)
	p.SetAge(19)
	p.SetAge(20)
}
