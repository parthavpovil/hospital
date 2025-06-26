package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct{
	DB *sql.DB
}

type Patient struct{
	Name string `json:"name"`
	Age int	`json:"age"`
	Gender string `json:"gender"`
	Contact string `json:"contact"`
	Diagnosis string `json:"diagnosis"`
	Prescription string `json:"prescription"`
}
func (a *App) patientAdd(w http.ResponseWriter, r *http.Request){
	data,err:=io.ReadAll(r.Body)
	if err!=nil{
		http.Error(w,"failed to ready request body",http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	var temp Patient
	json.Unmarshal(data,&temp)
	query:="INSERT INTO patients (name,age,gender,contact,diagnosis,prescription) VALUES($1,$2,$3,$4,$5,$6)"
	_,err1:=a.DB.Exec(query,temp.Name,temp.Age,temp.Gender,temp.Contact,temp.Diagnosis,temp.Prescription)
	if err1 != nil {
	log.Printf("DB insert error: %v", err1) 
	http.Error(w, "Failed to enter data to database", http.StatusInternalServerError)
	return
}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w,"patient added suceessfully ")
}

func main(){

	db,err:=sql.Open("postgres","user=parkar password=2002 dbname=patients sslmode=disable")
	if err!=nil{
		log.Fatal(err)
	}
	defer db.Close()

	r:=mux.NewRouter()

	app:=&App{DB:db}
	r.HandleFunc("/patients",app.patientAdd).Methods("POST")
	r.HandleFunc("/patients",app.patientGet).Methods("GET")
	http.ListenAndServe(":8000",r)
}