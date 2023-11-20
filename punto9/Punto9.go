package punto9

import (
	"TALLERJSON/students"
	"TALLERJSON/utils"
	"html/template"
	"net/http"
)

type Estudiante = students.Estudiante

type Punto9Result struct {
	EdadInf  int
	EdadSup  int
	Promedio float64
}

func Punto9(estudiantes []Estudiante) []Punto9Result {
	CantidadNotasPorRango := make([]float64, 10)
	PromediosPorRango := make([]float64, 10) // Crear un slice para los promedios en 10 rangos (0-9, 10-19, 20-29, etc.)
	var resultados []Punto9Result

	// Recorrer los estudiantes
	for _, estudiante := range estudiantes {
		promedioEstudiante := utils.CalcularPromedio(estudiante.Cursos) // Calcular el promedio de todas las notas de cada estudiante
		edad := estudiante.Edad

		// Determinar el rango de edad al que pertenece el estudiante ej: 23/10 = 2.3 -> posicion 2 del slice || 48/10 = 4.8 -> pos 4
		rango := edad / 10

		// Acumular el promedio en el rango correspondiente
		PromediosPorRango[rango] += promedioEstudiante
		CantidadNotasPorRango[rango] += 1
	}

	// Imprimir los resultados solo para los rangos con promedios
	for rango, totalNotas := range CantidadNotasPorRango {
		if totalNotas > 0 {
			PromediosPorRango[rango] /= totalNotas
			resultado := Punto9Result{
				EdadInf:  rango * 10,
				EdadSup:  (rango+1)*10 - 1,
				Promedio: PromediosPorRango[rango],
			}
			resultados = append(resultados, resultado)

		}
	}
	return resultados

}

func Punto9Handler(w http.ResponseWriter, r *http.Request, estudiantes []Estudiante) {
	tmpl, err := template.ParseFiles("templates/punto9.html", "templates/navbar.html")
	if err != nil {
		http.Error(w, "Error al analizar la plantilla", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, Punto9(estudiantes))
	if err != nil {
		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
		return
	}
}
