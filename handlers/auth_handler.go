package handlers

import (
	"database/sql"
	"encoding/json"
	"hospital/db"
	"hospital/models"
	"io"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func(a *App) Login_Handler(w http.ResponseWriter, r *http.Request){
	var temp models.User
	data,err:=io.ReadAll(r.Body)
	if err!=nil{
		http.Error(w,"Error reading the body",http.StatusInternalServerError)
		return
	}
	err1:=json.Unmarshal(data,&temp)
	if err1!=nil{
		http.Error(w,"error converting json to struct",http.StatusInternalServerError)
		return

	}
	retrived_pass,retrived_role,err2:=db.GetUser(a.DB,temp.Username)
	if err2==sql.ErrNoRows{
		http.Error(w,"User does not exit",http.StatusInternalServerError)
		return
	}else if err2!=nil{
		http.Error(w, "error retrieving data: "+err2.Error(), http.StatusInternalServerError)
		return
	}
	err3:=bcrypt.CompareHashAndPassword([]byte(retrived_pass),[]byte(temp.Password))
	if err3==nil{
		response:=map[string]string{
			"message":"Login successfull",
			"role":retrived_role,
		}
		json.NewEncoder(w).Encode(response)

	}else{
		http.Error(w,"invalid password or role",http.StatusInternalServerError)
		return
	}


}