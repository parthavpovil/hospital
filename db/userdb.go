package db

import (
	"database/sql"
	"hospital/models"

	"golang.org/x/crypto/bcrypt"
)

func GetUser(db *sql.DB,username string)(string,string,error){
	query:="SELECT password,role FROM users WHERE username=$1"
	var pass,role string
	err:=db.QueryRow(query,username).Scan(&pass,&role)
	return pass,role,err
}
func PostUser(db *sql.DB,p models.User)error{

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	query:="INSERT INTO users VALUES($1,$2,$3)"
	_,err1:=db.Exec(query,p.Username,hashedPassword,p.Role)
	if err1!=nil{
		return err
	}
	return nil

}