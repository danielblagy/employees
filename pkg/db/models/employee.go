package models

import "github.com/go-pg/pg/v10"

type Employee struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
}

func CreateEmployee(db *pg.DB, req *Employee) (*Employee, error) {
	_, err := db.Model(req).Insert()
	if err != nil {
		return nil, err
	}

	employee := &Employee{}

	err = db.Model(employee).Select()

	return employee, err
}

func GetEmployee(db *pg.DB, employeeID string) (*Employee, error) {
	employee := &Employee{}

	err := db.Model(employee).Select()

	return employee, err
}

func GetEmployees(db *pg.DB) ([]*Employee, error) {
	employees := make([]*Employee, 0)

	err := db.Model(&employees).Select()

	return employees, err
}
