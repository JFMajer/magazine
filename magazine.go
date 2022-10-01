package magazine

import (
	"fmt"
)

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	Address
}

type Employee struct {
	Name   string
	Salary float64
	Address
}

type Address struct {
	City       string
	Street     string
	PostalCode string
	Country    string
}

func ShowSub(sub *Subscriber) {
	if sub.Active {
		fmt.Printf("%s has status active with rate %.2f\n", sub.Name, sub.Rate)
	} else {
		fmt.Printf("%s has status inactive with rate %.2f\n", sub.Name, sub.Rate)
	}
}

func CreateSub(name string) *Subscriber {
	var newSub Subscriber
	newSub.Active = true
	newSub.Rate = 4.99
	newSub.Name = name

	return &newSub
}

func ApplyDiscount(s *Subscriber) {
	s.Rate = 2.55
}
