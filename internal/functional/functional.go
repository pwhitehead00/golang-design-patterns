// Functional Builder

package functional

import "fmt"

//Employee struct
type employee struct {
	name    string
	company string
	address string
}

type handler func(e *employee)

//EmployeeBuilder struct
type employeeBuilder struct {
	actions []handler
}

//Called sets name of the employee
func (b *employeeBuilder) Name(value string) *employeeBuilder {
	b.actions = append(b.actions, func(e *employee) {
		e.name = value
	})
	return b
}

//Company sets company of the employee
func (b *employeeBuilder) Company(value string) *employeeBuilder {
	b.actions = append(b.actions, func(e *employee) {
		e.company = value
	})
	return b
}

//At sets address of the employee
func (b *employeeBuilder) Address(value string) *employeeBuilder {
	b.actions = append(b.actions, func(e *employee) {
		e.address = value
	})
	return b
}

//Build builds the employee object
func (b *employeeBuilder) Build() employee {
	var emp employee
	for _, a := range b.actions {
		a(&emp)
	}
	return emp
}

//NewEmployeeBuilder - constructor
func NewEmployeeBuilder() *employeeBuilder {
	return &employeeBuilder{}
}

func (e *employee) Info() {
	fmt.Printf("Name: %s\nCompany: %s\nAddress: %s\n", e.name, e.company, e.address)
}
