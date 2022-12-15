package domain

type Employee struct {
	ID       string   `json:"id"`
	Name     string `json:"name"`
	Salary   uint   `json:"salary"`
	Position string `json:"position"`
}
