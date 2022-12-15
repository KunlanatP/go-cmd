package model

type CreateEmployeeDTO struct {
	Name     string `json:"name"`
	Salary   uint   `json:"salary"`
	Position string `json:"position"`
}

type UpdateEmployeeDTO struct {
	Salary   uint   `json:"salary"`
	Position string `json:"position"`
}
