package punto5

import (
	"TALLERJSON/students"
	"html/template"
	"net/http"
)

type Estudiante = students.Estudiante

type Punto5Result struct {
	Nombre string
	Edad   int
}

func Punto5(estudiantes []Estudiante) Punto5Result {
	var hombreMayorId int
	edad := 0
	for i := 0; i < len(estudiantes); i++ {
		if estudiantes[i].Gender == "male" {
			if estudiantes[i].Edad > edad {
				edad = estudiantes[i].Edad
				hombreMayorId = i
			}
		}
	}
	resultado := Punto5Result{
		Nombre: estudiantes[hombreMayorId].Nombre,
		Edad:   estudiantes[hombreMayorId].Edad,
	}
	return resultado
}

func Punto5Handler(w http.ResponseWriter, r *http.Request, estudiantes []Estudiante) {
	tmpl, err := template.ParseFiles("templates/punto5.html", "templates/navbar.html")
	if err != nil {
		http.Error(w, "Error al analizar la plantilla", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, Punto5(estudiantes))
	if err != nil {
		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
		return
	}
}
