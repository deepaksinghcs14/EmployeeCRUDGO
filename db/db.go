package db

import (
	"EmployeeCRUD/models"
)

var storage = make(map[string]models.Employee)

// Save stores a value in the storage
func Save(key string, value models.Employee) {
	storage[key] = value
}

func Update(key string, value models.Employee) models.Employee {
	employee := storage[key]
	if value.Name != "" {
		employee.Name = value.Name
	}
	if value.Position != "" {
		employee.Position = value.Position
	}
	if value.Salary != 0 {
		employee.Salary = value.Salary
	}
	storage[key] = employee
	return employee
}

func Retrieve(key string) models.Employee {
	return storage[key]
}

func Delete(key string) {
	delete(storage, key)
}

func List(page int, limit int) []models.Employee {
	// This is a simple pagination implementation
	// It is not efficient for large datasets
	if limit*(page-1) > len(storage) {
		return []models.Employee{}
	}
	startIndex := limit * (page - 1)
	endIndex := limit * page
	if endIndex > len(storage) {
		endIndex = len(storage)
	}
	employees := make([]models.Employee, 0, limit)
	for i := startIndex; i < endIndex; i++ {
		employees = append(employees, storage[string(i)])
	}
	return employees
}
