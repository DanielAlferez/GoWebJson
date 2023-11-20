package punto1

import (
	"TALLERJSON/students"
	"TALLERJSON/utils"
	"net/http"
	"text/template"
)

type Estudiante = students.Estudiante
type ResultadoPunto1 struct {
	Nombre        string
	Apellido      string
	MejorPromedio float64
}

// Punto1 retorna el nombre, apellido y el mejor promedio de un conjunto de estudiantes.
func Punto1(estudiantes []Estudiante) ResultadoPunto1 {
	resultado := ResultadoPunto1{
		Nombre:        estudiantes[0].Nombre,
		Apellido:      estudiantes[0].Apellido,
		MejorPromedio: utils.CalcularPromedio(estudiantes[0].Cursos),
	}

	for _, estudiante := range estudiantes {
		promedioActual := utils.CalcularPromedio(estudiante.Cursos)
		if promedioActual > resultado.MejorPromedio {
			resultado.MejorPromedio = promedioActual
			resultado.Nombre = estudiante.Nombre
			resultado.Apellido = estudiante.Apellido
		}
	}

	return resultado
}

func Punto1Handler(w http.ResponseWriter, r *http.Request, estudiante []Estudiante) {
	tmpl, err := template.ParseFiles("templates/punto1.html", "templates/navbar.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, Punto1(estudiante))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
