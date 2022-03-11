package main

import (
	"fmt"
	"log"
	"time"

	"github.com/pwhitehead00/golang-design-patterns/internal/constructor"
	"github.com/pwhitehead00/golang-design-patterns/internal/functional"
	"github.com/pwhitehead00/golang-design-patterns/internal/interfaces"
	optoins "github.com/pwhitehead00/golang-design-patterns/internal/options"
	"github.com/pwhitehead00/golang-design-patterns/internal/parameter"
)

func main() {

	// Basic Constructor pattern
	c := constructor.New("Paul", "Splunk", "123 Lincoln St")
	c.Info()

	// Constructor that takes a struct as an argument
	svr := constructor.NewServer(constructor.Config{
		Host:    "localhost",
		Port:    8080,
		Timeout: time.Duration(100) * time.Microsecond,
		MaxConn: 25,
	})
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}

	// Builder Pattern: Functional
	paul := functional.NewEmployeeBuilder().Name("Paul").Company("Splunk").Address("123 Lincoln St").Build()
	paul.Info()

	john := functional.NewEmployeeBuilder().Name("John").Company("Slack").Address("456 Forrest Way").Build()
	john.Info()

	// Builder Pattern: Parameters
	parameter.Info(func(b *parameter.EmployeeBuilder) {
		b.Name("Paul").Company("Splunk").Address("123 Lincoln st")
	})

	// Options or Functionall Options pattern.

	var opts []func(*optoins.Employee)
	opts = append(opts, optoins.WithCompany("foo"), optoins.WithAddress("123 Lincoln St"))

	o := optoins.New(opts...)
	o.Info()

	// Interfaces
	myGitHubInterface := interfaces.NewProject("github")
	fmt.Println(myGitHubInterface.GetRepo("my github repo"))

	myGitLabInterface := interfaces.NewProject("gitlab")
	fmt.Println(myGitLabInterface.GetRepo("my gitlab repo"))
}
