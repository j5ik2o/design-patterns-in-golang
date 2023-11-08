package observer

import "math/rand"

type RandomNumberGenerator struct {
	observers []Observer
	number    uint
}

func NewRandomNumberGenerator() *RandomNumberGenerator {
	return &RandomNumberGenerator{
		observers: []Observer{},
		number:    0,
	}
}

func (r *RandomNumberGenerator) AddObserver(observer Observer) {
	r.observers = append(
		r.observers,
		observer,
	)
}

func (r *RandomNumberGenerator) DeleteObserver(observer Observer) {
	var result []Observer
	for _, o := range r.observers {
		if o != observer {
			result = append(result, o)
		}
	}
	r.observers = result
}

func (r *RandomNumberGenerator) GetObservers() []Observer {
	return r.observers
}

func (r *RandomNumberGenerator) NotifyObservers() {
	for _, o := range r.observers {
		o.Update(r)
	}
}

func (r *RandomNumberGenerator) GetNumber() uint {
	return r.number
}

func (r *RandomNumberGenerator) Execute() {
	for i := 0; i < 20; i++ {
		r.number = uint(rand.Intn(50))
		r.NotifyObservers()
	}
}
