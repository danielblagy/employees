package api

import (
	"encoding/json"
	"log"
	"net/http"

	"employees/pkg/db/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-pg/pg/v10"
)

// start api with the pgdb and return a chi router
func StartAPI(pgdb *pg.DB) *chi.Mux {
	//get the router
	r := chi.NewRouter()
	//add middleware
	//in this case we will store our DB to use it later
	r.Use(middleware.Logger, middleware.WithValue("DB", pgdb))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to employees API!"))
	})

	//routes for our service
	r.Route("/employees", func(r chi.Router) {
		r.Post("/", createEmployee)
		r.Get("/", getEmployees)
	})

	return r
}

type EmployeeResponse struct {
	Success  bool             `json:"success"`
	Error    string           `json:"error"`
	Employee *models.Employee `json:"employee"`
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	//get the request body and decode it
	req := &models.Employee{}
	err := json.NewDecoder(r.Body).Decode(req)
	//if there's an error with decoding the information
	//send a response with an error
	if err != nil {
		res := &EmployeeResponse{
			Success:  false,
			Error:    err.Error(),
			Employee: nil,
		}
		err = json.NewEncoder(w).Encode(res)
		//if there's an error with encoding handle it
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		//return a bad request and exist the function
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//get the db from context
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	//if we can't get the db let's handle the error
	//and send an adequate response
	if !ok {
		res := &EmployeeResponse{
			Success:  false,
			Error:    "could not get the DB from context",
			Employee: nil,
		}
		err = json.NewEncoder(w).Encode(res)
		//if there's an error with encoding handle it
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		//return a bad request and exist the function
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//if we can get the db then
	employee, err := models.CreateEmployee(pgdb, req)
	if err != nil {
		res := &EmployeeResponse{
			Success:  false,
			Error:    err.Error(),
			Employee: nil,
		}
		err = json.NewEncoder(w).Encode(res)
		//if there's an error with encoding handle it
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		//return a bad request and exist the function
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//everything is good
	//let's return a positive response
	res := &EmployeeResponse{
		Success:  true,
		Error:    "",
		Employee: employee,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding after creating employee %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

type EmployeesResponse struct {
	Success   bool               `json:"success"`
	Error     string             `json:"error"`
	Employees []*models.Employee `json:"employees"`
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	//get db from ctx
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &EmployeesResponse{
			Success:   false,
			Error:     "could not get DB from context",
			Employees: nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//call models package to access the database and return the employees
	employees, err := models.GetEmployees(pgdb)
	if err != nil {
		res := &EmployeesResponse{
			Success:   false,
			Error:     err.Error(),
			Employees: nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//positive response
	res := &EmployeesResponse{
		Success:   true,
		Error:     "",
		Employees: employees,
	}
	//encode the positive response to json and send it back
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding employees: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
