package models

type Employee struct {
	ID       string  `json:"user_id,required"`
	Name     string  `json:"name,required"`
	Position string  `json:"position,required"`
	Salary   float64 `json:"salary,omitempty"`
}
