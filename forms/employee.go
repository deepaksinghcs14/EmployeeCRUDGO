package forms

type EmployeeSignUp struct {
	Name     string  `json:"name" binding:"required"`
	Position string  `json:"position" binding:"required"`
	Salary   float64 `json:"salary" binding:"required"`
}

type EmployeeUpdate struct {
	Name     string  `json:"name,omitempty"`
	Position string  `json:"position,omitempty"`
	Salary   float64 `json:"salary,omitempty"`
}

type Pagination struct {
	Page  int `json:"page" form:"page,omitempty,default=1"`
	Limit int `json:"limit" form:"limit,omitempty,default=10"`
}
