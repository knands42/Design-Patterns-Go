package observers

import (
	"fmt"
	"patterns/observer/observer_observable/observable"
)

type TrafficManagement struct {
	o observable.Observable
}

func NewTrafficManagement(o observable.Observable) *TrafficManagement {
	return &TrafficManagement{o: o}
}

func (t *TrafficManagement) Notify(data interface{}) {
	// Property Observer
	if val, ok := data.(observable.PropertyChangeEvent); ok {
		if val.Name == "age" {
			if val.Value.(int) >= 16 {
				fmt.Println("Congrats, you can drive now")
				t.o.Unsubscribe(t)
			}
		}
	}
}
