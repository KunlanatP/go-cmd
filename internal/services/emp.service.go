package services

import (
	"github.com/google/uuid"
	"github.com/kunlanat/go-cdm/internal/api/v1/employee/model"
	"github.com/kunlanat/go-cdm/internal/domain"
	"github.com/kunlanat/go-cdm/internal/errs"
	"github.com/kunlanat/go-cdm/internal/repository"
	"github.com/kunlanat/go-cdm/internal/repository/entities"
)

func New() EmpService {
	return &EmpSvc{
		repo: repository.New(),
	}
}

type EmpService interface {
	GetAllEmployees() (*[]entities.Employee, error)
	CreateEmployee(emp model.CreateEmployeeDTO) (*domain.Employee, error)
	UpdateEmployeeByID(id string, emp model.UpdateEmployeeDTO) (*domain.Employee, error)
}

type EmpSvc struct {
	repo repository.EmpRepository
}

func (r *EmpSvc) GetAllEmployees() (*[]entities.Employee, error) {
	return r.repo.GetAllEmployees()
}

func (r *EmpSvc) CreateEmployee(emp model.CreateEmployeeDTO) (*domain.Employee, error) {

	entity := entities.Employee{
		ID:       uuid.NewString(),
		Name:     emp.Name,
		Salary:   emp.Salary,
		Position: emp.Position,
	}
	return r.repo.CreateEmployee(entity)
}

func (r *EmpSvc) UpdateEmployeeByID(id string, emp model.UpdateEmployeeDTO) (*domain.Employee, error) {
	if len(id) < 36 {
		return nil, errs.ErrorIDFormat
	}

	entity := entities.Employee{
		Salary:   emp.Salary,
		Position: emp.Position,
	}

	return r.repo.UpdateEmployeeByID(id, entity)
}
