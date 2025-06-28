package db

import (
	"database/sql"
	"hospital/models"
)

func AddPatient(db *sql.DB, p models.Patient) error {
	query := `INSERT INTO patients (name, age, gender, contact, diagnosis, prescription) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(query, p.Name, p.Age, p.Gender, p.Contact, p.Diagnosis, p.Prescription)
	return err
}

func GetAllPatient(db *sql.DB)([]models.Patient,error){
	query:=`SELECT id,name,age, gender, contact, diagnosis, prescription FROM patients`
	rows,err:=db.Query(query)
	if err!=nil{
		return nil,err
	}
	defer rows.Close()

	var patients []models.Patient
	for rows.Next(){
		var p models.Patient
		err:=rows.Scan(&p.ID,&p.Name,&p.Age,&p.Gender,&p.Contact,&p.Diagnosis,&p.Prescription)
		if err!=nil{
			return nil,err
		}
		patients=append(patients, p)
	}
	return patients,nil
}

func UpdatePatientInfo(db *sql.DB,id int, temp models.Patient)error  {
	query:="UPDATE patients SET name=$1, age=$2, gender=$3, contact=$4 WHERE id=$5"
	_,err:=db.Exec(query,&temp.Name,&temp.Age,&temp.Gender,&temp.Contact,id)
	if err!=nil{
		return err
	}
	return nil
}

func UpdatePatientMedical(db *sql.DB,id int, temp models.Patient)error  {
	query:="UPDATE patients SET diagnosis=$1, prescription=$2 WHERE id=$3"
	_,err:=db.Exec(query,&temp.Diagnosis,&temp.Prescription,id)
	if err!=nil{
		return err
	}
	return nil
}

func DeletePatient(db *sql.DB,id int)error{
	query:="DELETE FROM patients WHERE id=$1"
	_,err:=db.Exec(query,id)
	if err!=nil{
		return err
	}
	return nil
}
func GetPatientById(db *sql.DB,id int)(models.Patient,error){
	query:="SELECT * FROM patients WHERE id=$1"
	var temp1 models.Patient
	err:=db.QueryRow(query,id).Scan(&temp1.ID,&temp1.Name,&temp1.Age,&temp1.Gender,&temp1.Contact,&temp1.Diagnosis,&temp1.Prescription)
	if err!=nil{
		return temp1 ,err
	}
	return temp1,nil

}