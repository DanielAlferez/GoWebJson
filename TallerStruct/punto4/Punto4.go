// punto4.go
package punto4

import (
	"TALLERJSON/students"
	"TALLERJSON/utils"
	"html/template"
	"net/http"
)

type Estudiante = students.Estudiante

type NombreNota struct {
	Nota   float64
	Nombre string
}

// Punto4Result contiene la información del curso y las notas.
type Punto4Result struct {
	Curso      string
	Top10Notas []NombreNota
}

// Punto4 devuelve el resultado del punto 4 como un struct Punto4Result.
func Punto4(estudiantes []Estudiante) []Punto4Result {
	var resultados []Punto4Result
	cursos := utils.ObtenerCursos(estudiantes)

	for _, curso := range cursos {
		top10 := obtenerPeores10Curso(estudiantes, curso)
		resultado := Punto4Result{
			Curso:      curso,
			Top10Notas: top10,
		}
		resultados = append(resultados, resultado)
	}

	return resultados
}

func obtenerPeores10Curso(estudiantes []Estudiante, nombreCurso string) []NombreNota {
	var notas []float64
	var nombresEstudiantes []string
	var notaTop float64
	var estudianteTop string
	var resultados []NombreNota

	for i := 0; i < 10; i++ {
		notaTop = 5
		for _, estudiante := range estudiantes {
			if utils.StringNoEstaEnSlice(estudiante.Nombre+" "+estudiante.Apellido, nombresEstudiantes) {
				for _, curso := range estudiante.Cursos {
					if curso.Curso == nombreCurso {
						if curso.Nota < notaTop {
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

// Punto4Handler es un manejador HTTP para el punto 4.
func Punto4Handler(w http.ResponseWriter, r *http.Request, estudiantes []Estudiante) {
	tmpl, err := template.ParseFiles("templates/punto4.html", "templates/navbar.html")
	if err != nil {
		http.Error(w, "Error al analizar la plantilla", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, Punto4(estudiantes))
	if err != nil {
		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
		return
	}
}
