package optoins

import (
	"fmt"
	"log"
)

type Employee struct {
	name    string
	company string
	address string
}

func New(options ...func(*Employee)) *Employee {
	e := new(Employee)
	for _, o := range options {
		o(e)
	}

	if e.name == "" {
		e.name = "John"
	}

	if e.address == "" {
		log.Fatal("Address must be specified")
	}
	return e
}

func (e *Employee) Info() {
	fmt.Printf("Name: %s\nCompany: %s\nAddress: %s\n", e.name, e.company, e.address)
}

func WithName(name string) func(*Employee) {
	return func(e *Employee) {
		e.name = name
	}
}

func WithCompany(company string) func(*Employee) {
	return func(e *Employee) {
		e.company = company
	}
}

func WithAddress(address string) func(*Employee) {
	return func(e *Employee) {
		e.address = address
	}
}
