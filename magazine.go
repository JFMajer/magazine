package magazine

import (
	"fmt"
)

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
}

func showSub(sub *Subscriber) {
	if sub.Active {
		fmt.Printf("%s has status active with rate %.2f\n", sub.Name, sub.Rate)
	} else {
		fmt.Printf("%s has status inactive with rate %.2f\n", sub.Name, sub.Rate)
	}
}

func createSub(name string) *Subscriber {
	var newSub Subscriber
	newSub.Active = true
	newSub.Rate = 4.99
	newSub.Name = name

	return &newSub
}

func applyDiscount(s *Subscriber) {
	s.Rate = 2.55
}
