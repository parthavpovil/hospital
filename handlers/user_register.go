package handlers

import (
	"encoding/json"
	"fmt"
	"hospital/db"
	"hospital/models"
	"io"
	"net/http"
)

func (a *App)User_Signup(w http.ResponseWriter, r *http.Request){
data,err:=io.ReadAll(r.Body)
if err!=nil{
	http.Error(w,"error reading body ",http.StatusInternalServerError)
	return
}


var temp models.User
err1:=json.Unmarshal(data,&temp)
if err1!=nil{
	http.Error(w,"error converting from json",http.StatusInternalServerError)
	return
}
err2:=db.PostUser(a.DB,temp)
if err2!=nil{
	http.Error(w,"error adding user",http.StatusInternalServerError)
	return
}
w.WriteHeader(http.StatusOK)
fmt.Fprintln(w,"User successfully added")
}