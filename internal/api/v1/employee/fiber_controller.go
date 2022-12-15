package employee

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kunlanat/go-cdm/internal/api/v1/employee/model"
	"github.com/kunlanat/go-cdm/internal/services"
)

func RegisterRoute(route fiber.Router) {
	empCtl := empCtl{svc: services.New()}
	route.Get("/fiber/employees", empCtl.GetAllEmployees)
	route.Post("/fiber/employees", empCtl.CreateEmployee)
	route.Patch("/fiber/employees/:id", empCtl.UpdateEmployeeByID)
}

type empCtl struct {
	svc services.EmpService
}

func (s *empCtl) GetAllEmployees(c *fiber.Ctx) error {

	data, err := s.svc.GetAllEmployees()
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(data)
}

func (s *empCtl) CreateEmployee(c *fiber.Ctx) error {
	emp := model.CreateEmployeeDTO{}
	if err := c.BodyParser(&emp); err != nil {
		return err
	}

	data, err := s.svc.CreateEmployee(emp)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(data)
}

func (s *empCtl) UpdateEmployeeByID(c *fiber.Ctx) error {
	empId := c.Params("id")
	emp := model.UpdateEmployeeDTO{}

	if err := c.BodyParser(&emp); err != nil {
		return err
	}

	data, err := s.svc.UpdateEmployeeByID(empId, emp)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(data)
}
