// Example - Builder Parameter

package parameter

import (
	"fmt"
)

//employee struct
type employee struct {
	name    string
	company string
	address string
}

// Generic private function to build the employee
type action func(builder *EmployeeBuilder)

//EmployeeBuilder struct, public struct
type EmployeeBuilder struct {
	employee employee
}

//Send function is for client to send email
func Info(a action) {
	var builder EmployeeBuilder
	a(&builder)
	fmt.Printf("Name: %s\nCompany: %s\nAddress: %s\n", builder.employee.name, builder.employee.company, builder.employee.address)
}

//At sets address of the employee
func (eb *EmployeeBuilder) Address(value string) *EmployeeBuilder {
	eb.employee.address = value
	return eb
}

//Company sets company of the employee
func (eb *EmployeeBuilder) Company(value string) *EmployeeBuilder {
	eb.employee.company = value
	return eb
}

//Called sets name of the employee
func (eb *EmployeeBuilder) Name(value string) *EmployeeBuilder {
	eb.employee.name = value
	return eb
}
