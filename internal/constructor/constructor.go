// Constructor Pattern

package constructor

import (
	"fmt"
	"time"
)

type employee struct {
	name    string
	company string
	address string
}

func New(name, company, address string) *employee {
	return &employee{
		name:    name,
		company: company,
		address: address,
	}
}
func (e *employee) Info() {
	fmt.Printf("Name: %s\nCompany: %s\nAddress: %s\n", e.name, e.company, e.address)
}

type Server struct {
	cfg Config
}

type Config struct {
	Host    string
	Port    int
	Timeout time.Duration
	MaxConn int
}

func NewServer(cfg Config) *Server {
	return &Server{cfg}
}

func (s *Server) Start() error {
	fmt.Println("Starting server ...")
	return nil
}
