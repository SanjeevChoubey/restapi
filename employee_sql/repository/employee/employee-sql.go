package repository

import (
	"database/sql"
	"log"

	"github.com/SanjeevChoubey/RESTapi/employee_sql/models"
)

type Repository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func (r *Repository) GetEmployees(db *sql.DB, employee models.Employee, employees []models.Employee) []models.Employee {

	rows, err := db.Query("Select * from employee")
	logFatal(err)
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&employee.ID, &employee.Name, &employee.Age, &employee.Sex, &employee.City, &employee.State, &employee.Country)
		logFatal(err)
		employees = append(employees, employee)
	}
	return employees
}

func (r *Repository) GetEmployee(db *sql.DB, employee models.Employee, id int) models.Employee {

	err := db.QueryRow("Select * from employee where id = ?", id).Scan(&employee.ID, &employee.Name, &employee.Age, &employee.Sex, &employee.City, &employee.State, &employee.Country)
	if err != nil {
		// if no row in db for given query then donot do any thing
		if err != sql.ErrNoRows {
			log.Fatal(err)
		}
	}

	return employee
}

func (r *Repository) AddEmployee(db *sql.DB, employee models.Employee) int64 {
	res,err:= db.Exec("Insert into employee(id,name,age,sex,city,state,country) values(?,?,?,?,?,?,?) ",
	&employee.ID, &employee.Name, &employee.Age, &employee.Sex, &employee.City, &employee.State, &employee.Country)
	logFatal(err)
	rowsaffected,err:= res.RowsAffected()
	logFatal(err)
	return rowsaffected
}

func (r *Repository) UpdateEmployee(db *sql.DB, employee models.Employee) int64 {
	res,err:= db.Exec("Update employee set name =?,age=?,sex=?,city=?,state=?,country=? where id =?", 
	&employee.Name, &employee.Age, &employee.Sex, &employee.City, &employee.State, &employee.Country,&employee.ID)
	logFatal(err)
	rowsaffected,err:= res.RowsAffected()
	logFatal(err)
	return rowsaffected
}

func (r *Repository) DeleteEmployee(db *sql.DB, id int) int64 {
	res,err:= db.Exec("Delete from employee where id =?", id)
	logFatal(err)
	rowsaffected,err:= res.RowsAffected()
	logFatal(err)
	return rowsaffected
}