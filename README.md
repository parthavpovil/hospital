# Hospital Management Backend (Golang)

This is a simple backend API for a hospital management system, built with Go. It supports both receptionist and doctor roles, with JWT-based authentication and role-based access to patient data.

---

## Features

- **Single login** for both doctors and receptionists
- **Receptionist**: Register new patients, view, update, and delete patient records
- **Doctor**: View all patients, update diagnosis and prescription
- **Role-based access control** using JWT
- **RESTful API** structure

---

## Prerequisites

- Go (I used version 1.20+, but any recent version should work)
- PostgreSQL (make sure it's running locally or update the config for your setup)
- Postman (for testing the API, optional but recommended)

---

## Getting Started

### 1. Clone the Repository

```sh
git clone https://github.com/parthavpovil/hospital.git
cd hospital
```

### 2. Set Up the Database

- Create a PostgreSQL database (I named mine `patients`).
- Create the required tables. Here's a quick SQL you can use:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL
);

CREATE TABLE patients (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    age INT,
    gender VARCHAR(10),
    contact VARCHAR(20),
    diagnosis TEXT,
    prescription TEXT
);
```

- You can run this in pgAdmin or with `psql`.

### 3. Configure Database Connection

- The DB connection string is in `main.go`. Update the user, password, and dbname as per your local setup:
  ```go
  sql.Open("postgres", "user=YOURUSER password=YOURPASS dbname=patients sslmode=disable")
  ```

### 4. Install Dependencies

```sh
go mod tidy
```

### 5. Run the Application

```sh
go run main.go
```

---

## API Documentation

I used Postman to test the API. Here are the main endpoints:

| Method | Endpoint                        | Description                                 | Auth Required |
|--------|---------------------------------|---------------------------------------------|--------------|
| POST   | /login                          | Login for both roles                        | No           |
| POST   | /signup                         | Register a new user                         | No           |
| GET    | /patients/                      | Get all patients                            | Yes          |
| POST   | /patients/                      | Add a new patient                           | Yes          |
| GET    | /patients/{id}                  | Get patient by ID                           | Yes          |
| PUT    | /patients/reception/{id}        | Receptionist updates patient info           | Yes          |
| PUT    | /patients/doctor/{id}           | Doctor updates diagnosis/prescription       | Yes          |
| DELETE | /patients/{id}                  | Delete patient                              | Yes          |

- For protected routes, add an `Authorization` header with the JWT token you get from `/login`.

### Postman Collection

You can use my Postman collection to try out all the endpoints easily:
[Open Hospital API Postman Collection](https://parkar-1071016.postman.co/workspace/local-api~cba269b2-55f3-42e9-887f-b0060ed8576b/collection/46028416-57a02136-77c8-4cfe-b50e-1c18f6552940?action=share&creator=46028416)

**Instructions:**
- Import the collection into Postman using the link above.
- For requests that need authentication, log in first and copy the token from the login response.
- Add your token in the `Authorization` header for protected routes (just paste the token, no "Bearer" needed).
- When deleting or updating a patient, replace `{id}` in the URL with the actual patient ID you want to use.
- For POST and PUT requests, set the body type to `raw` and select `JSON`.
- Make sure your backend server is running before you try these requests.

### Example: Update Patient as Doctor

- **PUT** `/patients/doctor/1`
- **Headers:** `Authorization: <your_token>`, `Content-Type: application/json`
- **Body:**
  ```json
  {
    "diagnosis": "Updated diagnosis",
    "prescription": "Updated prescription"
  }
  ```

---

## Running Tests

To run unit tests (if present):

```sh
go test ./...
```

---

## Notes

- If you want to use different DB credentials, update them in `main.go`.
- If you want to add more features, feel free! (I kept it simple for the assignment.)
- If you have any issues, check your DB connection and make sure the tables exist.

---

## Contact

If you have questions or suggestions, feel free to open an issue or contact me.

---

**Tip:**  
If you want to try the API quickly, import the endpoints into Postman and use the example requests above.

---

Let me know if you want to add anything else or need a sample Postman collection!