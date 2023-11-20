package punto6

import (
	"TALLERJSON/students"
	"html/template"
	"net/http"
)

type Estudiante = students.Estudiante

type Punto6Result struct {
	Nombre string
	Edad   int
}

func Punto6(estudiantes []Estudiante) Punto6Result {
	var mujerMayorId int
	edad := 0
	for i := 0; i < len(estudiantes); i++ {
		if estudiantes[i].Gender == "female" {
			if estudiantes[i].Edad > edad {
				edad = estudiantes[i].Edad
				mujerMayorId = i
			}
		}
	}

	resultado := Punto6Result{
		Nombre: estudiantes[mujerMayorId].Nombre,
		Edad:   estudiantes[mujerMayorId].Edad,
	}
	return resultado
}

func Punto6Handler(w http.ResponseWriter, r *http.Request, estudiantes []Estudiante) {
	tmpl, err := template.ParseFiles("templates/punto6.html", "templates/navbar.html")
	if err != nil {
		http.Error(w, "Error al analizar la plantilla", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, Punto6(estudiantes))
	if err != nil {
		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
		return
	}
}
