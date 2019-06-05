package models

type Employee struct {
	ID      int    `json:id`
	Name    string `json:name`
	Age     int    `json:age`
	Sex     string `json:sex`
	City    string `json:city`
	State   string `json:state`
	Country string `json:country`
}

var Employees []Employee
