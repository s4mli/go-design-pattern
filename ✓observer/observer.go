package observer

import "math/rand"

type numberObserver interface {
	update() int
}

type numberGenerator struct {
	observers []numberObserver
}

func (ng *numberGenerator) AddObserver(o numberObserver) {
	ng.observers = append(ng.observers, o)
}

func (ng *numberGenerator) notifyObservers() []int {
	results := []int{}
	for _, ob := range ng.observers {
		results = append(results, ob.update())
	}
	return results
}

/******************************************************************/

type DigitObserver struct {
}

func (do *DigitObserver) update() int {
	return rand.Intn(100)
}

/******************************************************************/

type randomNumberGenerator struct {
	*numberGenerator
}

func (r *randomNumberGenerator) Execute() []int {
	return r.notifyObservers()
}

func NewRandomNumberGenerator() *randomNumberGenerator {
	return &randomNumberGenerator{&numberGenerator{}}
}
