package main

import (
	"TALLERJSON/punto1"
	"TALLERJSON/punto2"
	"TALLERJSON/punto3"
	"TALLERJSON/punto4"
	"TALLERJSON/punto5"
	"TALLERJSON/punto6"
	"TALLERJSON/punto7"
	"TALLERJSON/punto8"
	"TALLERJSON/punto9"
	"TALLERJSON/students"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type estudiante = students.Estudiante

func main() {
	file, err := os.Open("generated.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	var estudiantes []estudiante
	err = json.Unmarshal(data, &estudiantes)
	if err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/punto2", func(w http.ResponseWriter, r *http.Request) {
		punto2.Punto2Handler(w, r, estudiantes)
	})
	http.HandleFunc("/punto4", func(w http.ResponseWriter, r *http.Request) {
		punto4.Punto4Handler(w, r, estudiantes)
	})
	http.HandleFunc("/punto6", func(w http.ResponseWriter, r *http.Request) {
		punto6.Punto6Handler(w, r, estudiantes)
	})
	http.HandleFunc("/punto8", func(w http.ResponseWriter, r *http.Request) {
		punto8.Punto8Handler(w, r, estudiantes)
	})
	http.HandleFunc("/punto1", func(w http.ResponseWriter, r *http.Request) {
		punto1.Punto1Handler(w, r, estudiantes)
	})
	http.HandleFunc("/punto3", func(w http.ResponseWriter, r *http.Request) {
		punto3.Punto3Handler(w, r, estudiantes)
	})
	http.HandleFunc("/punto5", func(w http.ResponseWriter, r *http.Request) {
		punto5.Punto5Handler(w, r, estudiantes)
	})
	http.HandleFunc("/punto7", func(w http.ResponseWriter, r *http.Request) {
		punto7.Punto7Handler(w, r, estudiantes)
	})
	http.HandleFunc("/punto9", func(w http.ResponseWriter, r *http.Request) {
		punto9.Punto9Handler(w, r, estudiantes)
	})
	
	http.ListenAndServe("localhost:8090", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home.html", "templates/navbar.html")
	if err != nil {
		http.Error(w, "Error al analizar la plantilla", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
		return
	}
}
