package punto3

import (
	"TALLERJSON/students"
	"TALLERJSON/utils"
	"net/http"
	"text/template"
)

type Estudiante = students.Estudiante

type NombreNota struct {
	Nota   float64
	Nombre string
}

// Punto3Result contiene la información del curso y las notas.
type Punto3Result struct {
	Curso      string
	Top10Notas []NombreNota
}

func Punto3(estudiantes []Estudiante) []Punto3Result {
	cursos := utils.ObtenerCursos(estudiantes)
	var resultados []Punto3Result

	for _, curso := range cursos {

		resultadoCurso := obtenerMejores10Curso(estudiantes, curso)
		resultado := Punto3Result{
			Curso:      curso,
			Top10Notas: resultadoCurso,
		}
		resultados = append(resultados, resultado)
	}

	return resultados
}

// obtenerMejores10Curso retorna una estructura con las notas y nombres de los 10 mejores estudiantes en un curso.
func obtenerMejores10Curso(estudiantes []Estudiante, nombreCurso string) []NombreNota {
	var resultados []NombreNota
	var notas []float64
	var nombresEstudiantes []string

	for i := 0; i < 10; i++ {
		notaTop := 0.0
		estudianteTop := ""

		for _, estudiante := range estudiantes {
			if utils.StringNoEstaEnSlice(estudiante.Nombre+" "+estudiante.Apellido, nombresEstudiantes) {
				for _, curso := range estudiante.Cursos {
					if curso.Curso == nombreCurso {
						if curso.Nota > notaTop {
							estudianteTop = estudiante.Nombre + " " + estudiante.Apellido
							notaTop = curso.Nota
						}
					}
				}
			}
		}
		resultado := NombreNota{
			Nota:   notaTop,
			Nombre: estudianteTop,
		}
		resultados = append(resultados, resultado)
		notas = append(notas, notaTop)
		nombresEstudiantes = append(nombresEstudiantes, estudianteTop)
	}

	return resultados
}

func Punto3Handler(w http.ResponseWriter, r *http.Request, estudiante []Estudiante) {
	tmpl, err := template.ParseFiles("templates/punto3.html", "templates/navbar.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, Punto3(estudiante))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
