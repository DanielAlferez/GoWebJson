package punto8

import (
	"TALLERJSON/students"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"
)

type Estudiante = students.Estudiante

func Punto8(estudiantes []Estudiante) []string {
	var nombres2022 []string

	for _, estudiante := range estudiantes {
		fechaMatriculacion := strings.Replace(estudiante.Matriculado, " ", "", -1)

		matriculacion, err := time.Parse(time.RFC3339, fechaMatriculacion)
		if err != nil {
			fmt.Printf("Error al parsear fecha para %s: %v\n", estudiante.Nombre, err)
			continue
		}

		if matriculacion.Year() == 2022 {
			nombres2022 = append(nombres2022, estudiante.Nombre)
		}
	}

	return nombres2022

}

func Punto8Handler(w http.ResponseWriter, r *http.Request, estudiantes []Estudiante) {
	tmpl, err := template.ParseFiles("templates/punto8.html", "templates/navbar.html")
	if err != nil {
		http.Error(w, "Error al analizar la plantilla", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, Punto8(estudiantes))
	if err != nil {
		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
		return
	}
}
