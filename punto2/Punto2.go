// punto2.go
package punto2

import (
	"TALLERJSON/students"
	"TALLERJSON/utils"
	"html/template"
	"net/http"
)

type Estudiante = students.Estudiante

// Punto2Result contiene la información del estudiante y su promedio.
type Punto2Result struct {
	Estudiante Estudiante
	Promedio   float64
}

// Punto2 devuelve el resultado del punto 2 como un struct Punto2Result.
func Punto2(estudiantes []Estudiante) Punto2Result {
	peorPromedio := estudiantes[0]
	for _, estudiante := range estudiantes {
		if utils.CalcularPromedio(estudiante.Cursos) < utils.CalcularPromedio(peorPromedio.Cursos) {
			peorPromedio = estudiante
		}
	}

	promedio := utils.CalcularPromedio(peorPromedio.Cursos)
	resultado := Punto2Result{
		Estudiante: peorPromedio,
		Promedio:   promedio,
	}

	return resultado
}

func Punto2Handler(w http.ResponseWriter, r *http.Request, estudiantes []Estudiante) {
	tmpl, err := template.ParseFiles("templates/punto2.html", "templates/navbar.html")
	if err != nil {
		http.Error(w, "Error al analizar la plantilla", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, Punto2(estudiantes))
	if err != nil {
		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
		return
	}
}
