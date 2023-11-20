package punto7

import (
	"TALLERJSON/students"
	"html/template"
	"math"
	"net/http"
)

type Estudiante = students.Estudiante

type Punto7Result struct {
	Curso    string
	Promedio float64
	Rango    float64
	Varianza float64
	DE       float64
}

func calcularPromedio(numeros []float64) float64 {
	suma := 0.0
	for _, numero := range numeros {
		suma += numero
	}
	return suma / float64(len(numeros))
}

func calcularRango(numeros []float64) float64 {
	min := numeros[0]
	max := numeros[0]
	for _, numero := range numeros {
		if numero < min {
			min = numero
		}
		if numero > max {
			max = numero
		}
	}
	return max - min
}

func calcularVarianza(numeros []float64, promedio float64) float64 {
	sumaCuadrados := 0.0
	for _, numero := range numeros {
		sumaCuadrados += math.Pow(numero-promedio, 2)
	}
	return sumaCuadrados / float64(len(numeros))
}

func calcularDesviacionEstandar(varianza float64) float64 {
	return math.Sqrt(varianza)
}

func Punto7(estudiantes []Estudiante) []Punto7Result {
	var notas [][]float64
	var resultados []Punto7Result

	for i := 0; i < len(estudiantes[0].Cursos); i++ {
		var curso []float64
		for j := 0; j < len(estudiantes); j++ {
			curso = append(curso, estudiantes[j].Cursos[i].Nota)
		}
		notas = append(notas, curso)
	}

	for i := 0; i < len(notas); i++ {

		promedio := calcularPromedio(notas[i])
		rango := calcularRango(notas[i])
		varianza := calcularVarianza(notas[i], promedio)
		desviacion := calcularDesviacionEstandar(varianza)
		resultado := Punto7Result{
			Curso:    estudiantes[0].Cursos[i].Curso,
			Promedio: promedio,
			Rango:    rango,
			Varianza: varianza,
			DE:       desviacion,
		}
		resultados = append(resultados, resultado)

	}
	return resultados
}

func Punto7Handler(w http.ResponseWriter, r *http.Request, estudiantes []Estudiante) {
	tmpl, err := template.ParseFiles("templates/punto7.html", "templates/navbar.html")
	if err != nil {
		http.Error(w, "Error al analizar la plantilla", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, Punto7(estudiantes))
	if err != nil {
		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
		return
	}
}
