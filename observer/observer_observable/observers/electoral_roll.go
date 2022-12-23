package observers

import "patterns/observer/observer_observable/observable"

type ElectoralRoll struct {
	o observable.Observable
}

func NewElectoralRoll(o observable.Observable) *ElectoralRoll {
	return &ElectoralRoll{o: o}
}

func (e *ElectoralRoll) Notify(data interface{}) {
	// Property Observer
	if val, ok := data.(observable.PropertyChangeEvent); ok {
		if val.Name == "canVote" {
			if val.Value.(bool) {
				println("Congrats, you can vote now")
				e.o.Unsubscribe(e)
			}
		}
	}
}
