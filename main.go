package main

import (
	"log"
	"net/http"

	"database/sql"

	"hospital/handlers"
	"hospital/middleware"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)




func main(){

	db,err:=sql.Open("postgres","user=parkar password=2002 dbname=patients sslmode=disable")
	if err!=nil{
		log.Fatal(err)
	}
	defer db.Close()

	r:=mux.NewRouter()
	app:=&handlers.App{DB:db}

	r.HandleFunc("/login",app.Login_Handler).Methods("POST")
	r.HandleFunc("/signup",app.User_Signup).Methods("POST")

	secured := r.PathPrefix("/patients").Subrouter()
	secured.Use(middleware.AuthMiddleware)

	secured.HandleFunc("/", app.PatientAdd).Methods("POST")
	secured.HandleFunc("/", app.PatientGet).Methods("GET")
	secured.HandleFunc("/reception/{id}", app.PatientUpdateReception).Methods("PUT")
	secured.HandleFunc("/doctor/{id}", app.PatientUpdateDoctor).Methods("PUT")
	secured.HandleFunc("/{id}", app.PatientDelete).Methods("DELETE")
	secured.HandleFunc("/{id}",app.PatientGetbyId).Methods("GET")

	
	http.ListenAndServe(":8000",r)
}