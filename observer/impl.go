package observer

import (
	"fmt"
	"math/rand"
	"time"
)

type NumberGenerator interface {
	AddObserver(observer Observer)
	DeleteObserver(observer Observer)
	NotifyObservers()
	GetNumber() uint
	Execute()
}

type Observer interface {
	Update(generator NumberGenerator)
}

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

type DigitObserver struct {
}

func (d DigitObserver) Update(generator NumberGenerator) {
	fmt.Printf("DigitObserver:%d\n", generator.GetNumber())
	time.Sleep(100 * time.Millisecond)
}

type GraphObserver struct {
}

func (g GraphObserver) Update(generator NumberGenerator) {
	fmt.Print("GraphObserver:")
	count := generator.GetNumber()
	for i := uint(0); i < count; i++ {
		fmt.Print("*")
	}
	fmt.Println("")
	time.Sleep(100 * time.Millisecond)
}
