package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Alumno struct {
	Nombre    string `json:"nombre"`
	Matricula string `json:"matricula"`
}

type Profesor struct {
	Nombre         string `json:"nombre"`
	NumeroEmpleado string `json:"numeroEmpleado"`
}

func main() {
	alumnos := []Alumno{
		{"Juan Perez", "20230011"},
		{"Juan Gonzalez", "10101010"},
		{"Juan Perez", "20230011"},

	}

	profesores := []Profesor{
		{"Ing. Rodriguez", "00001"},
		{"Dr. Miranda", "00301"},
		{"Dr. Lopez", "00001"},
	}

	http.HandleFunc("/alumnos", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(alumnos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/profesores", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(profesores)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
