package main

import (
	
	"log"
	"net/http"

	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"hospital/handlers"
)




func main(){

	db,err:=sql.Open("postgres","user=parkar password=2002 dbname=patients sslmode=disable")
	if err!=nil{
		log.Fatal(err)
	}
	defer db.Close()

	r:=mux.NewRouter()

	app:=&handlers.App{DB:db}
	r.HandleFunc("/patients",app.PatientAdd).Methods("POST")
	r.HandleFunc("/patients",app.PatientGet).Methods("GET")
	r.HandleFunc("/patients/{id}",app.PatientUpdate).Methods("PUT")
	r.HandleFunc("/patients/{id}",app.PatientDelete).Methods("DELETE")
	r.HandleFunc("/login",app.Login_Handler).Methods("POST")
	r.HandleFunc("/signup",app.User_Signup).Methods("POST")
	http.ListenAndServe(":8000",r)
}