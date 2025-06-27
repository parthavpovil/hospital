package models

type Patient struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Age int	`json:"age"`
	Gender string `json:"gender"`
	Contact string `json:"contact"`
	Diagnosis string `json:"diagnosis"`
	Prescription string `json:"prescription"`
}

type User struct{
	Username string `json:"username"`
	Password string `json:"password"`
	Role string `json:"role"`
}