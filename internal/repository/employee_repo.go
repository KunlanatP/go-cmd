package repository

import (
	"github.com/google/uuid"
	"github.com/kunlanat/go-cdm/internal/domain"
	"github.com/kunlanat/go-cdm/internal/repository/entities"
)

func New() EmpRepository {
	return &EmpRepo{
		emps: []entities.Employee{
			{
				ID:       uuid.NewString(),
				Name:     "Harry",
				Salary:   2000,
				Position: "wizard",
			},
		},
	}
}

type EmpRepository interface {
	GetAllEmployees() (*[]entities.Employee, error)
	CreateEmployee(emp entities.Employee) (*domain.Employee, error)
	UpdateEmployeeByID(id string, emp entities.Employee) (*domain.Employee, error)
}

type EmpRepo struct {
	emps []entities.Employee
}

func (r *EmpRepo) GetAllEmployees() (*[]entities.Employee, error) {
	return &r.emps, nil
}

func (r *EmpRepo) CreateEmployee(emp entities.Employee) (*domain.Employee, error) {
	r.emps = append(r.emps, emp)
	return emp.ToDomain(), nil
}

func (r *EmpRepo) UpdateEmployeeByID(id string, emp entities.Employee) (*domain.Employee, error) {
	for index, employee := range r.emps {
		if employee.ID == id {
			employee.Salary = emp.Salary
			employee.Position = emp.Position

			r.emps[index] = employee

			return employee.ToDomain(), nil
		}
	}
	return nil, nil
}
