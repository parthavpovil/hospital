package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"hospital/db"
	"hospital/models"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	DB *sql.DB
}

func (a *App) PatientDelete(w http.ResponseWriter, r *http.Request) {
	iddata := mux.Vars(r)["id"]
	id, err := strconv.Atoi(iddata)
	if err != nil {
		http.Error(w, "error converting id", http.StatusInternalServerError)
		return
	}
	err1 := db.DeletePatient(a.DB, id)
	if err1 != nil {
		http.Error(w, "Error Deleting patient from db", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "patient deleted suceessfully ")

}

func (a *App) PatientUpdateReception(w http.ResponseWriter, r *http.Request) {
	role, ok := r.Context().Value("role").(string)
	if !ok {
		http.Error(w, "error getting role from context", http.StatusInternalServerError)
		return
	}
	if role != "reception" {
		http.Error(w, "you are not authorized to update patient personal details", http.StatusForbidden)
		return
	}
	iddata := mux.Vars(r)["id"]
	id, err := strconv.Atoi(iddata)
	if err != nil {
		http.Error(w, "error converting id to string", http.StatusInternalServerError)
		return
	}
	bodydata, err1 := io.ReadAll(r.Body)
	if err1 != nil {
		http.Error(w, "error reading body", http.StatusInternalServerError)
		return
	}
	var temp models.Patient
	err2 := json.Unmarshal(bodydata, &temp)
	if err2 != nil {
		http.Error(w, "error converting json to struct ", http.StatusInternalServerError)
		return
	}
	err3 := db.UpdatePatientInfo(a.DB, id, temp)
	if err3 != nil {
		http.Error(w, "Error upating patient", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "patient details updated suceessfully ")

}

func (a *App) PatientUpdateDoctor(w http.ResponseWriter, r *http.Request) {
	iddata := mux.Vars(r)["id"]
	iddata = strings.TrimSpace(iddata)
	//fmt.Println("iddata:", iddata)
	id, err := strconv.Atoi(iddata)
	if err != nil {
		http.Error(w, "error converting id to int", http.StatusInternalServerError)
		return
	}

	role, ok := r.Context().Value("role").(string)
	if !ok {
		http.Error(w, "error getting role from context", http.StatusInternalServerError)
		return
	}
	if role != "doctor" {
		http.Error(w, "you are not authorized to update patient medical details", http.StatusForbidden)
		return
	}
	
	bodydata, err1 := io.ReadAll(r.Body)
	if err1 != nil {
		http.Error(w, "error reading body", http.StatusInternalServerError)
		return
	}
	var temp models.Patient
	err2 := json.Unmarshal(bodydata, &temp)
	if err2 != nil {
		http.Error(w, "error converting json to struct ", http.StatusInternalServerError)
		return
	}
	err3 := db.UpdatePatientMedical(a.DB, id, temp)
	if err3 != nil {
		http.Error(w, "Error upating patient", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "patient details updated suceessfully ")

}

func (a *App) PatientAdd(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to ready request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	var temp models.Patient
	json.Unmarshal(data, &temp)
	err1 := db.AddPatient(a.DB, temp)
	if err1 != nil {
		log.Printf("DB insert error: %v", err1)
		http.Error(w, "Failed to enter data to database", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "patient added suceessfully ")
}

func (a *App) PatientGet(w http.ResponseWriter, r *http.Request) {
	temp, err := db.GetAllPatient(a.DB)
	if err != nil {
		http.Error(w, "error loading data from db", http.StatusInternalServerError)
		return
	}
	data, err1 := json.Marshal(&temp)
	if err1 != nil {
		http.Error(w, "error convert to json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

func (a *App) PatientGetbyId(w http.ResponseWriter, r *http.Request) {
	data := mux.Vars(r)["id"]
	id, err := strconv.Atoi(data)
	if err != nil {
		http.Error(w, "error convert id ", http.StatusInternalServerError)
		return
	}
	temp, err1 := db.GetPatientById(a.DB, id)
	if err1 != nil {
		http.Error(w, "error loading data from db", http.StatusInternalServerError)
		return
	}
	data1, err2 := json.Marshal(&temp)
	if err2 != nil {
		http.Error(w, "error convert to json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data1)

}
