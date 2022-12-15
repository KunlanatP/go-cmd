package entities

import (
	"github.com/kunlanat/go-cdm/internal/domain"
)

type Employee struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Salary   uint   `json:"salary"`
	Position string `json:"position"`
}

func (emp *Employee) ToDomain() *domain.Employee {
	return &domain.Employee{
		ID:       emp.ID,
		Name:     emp.Name,
		Salary:   emp.Salary,
		Position: emp.Position,
	}
}
