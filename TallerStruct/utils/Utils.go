package utils

import "TALLERJSON/students"

type Curso = students.Curso
type Estudiante = students.Estudiante

func CalcularPromedio(cursos []Curso) float64 {
	suma := 0.0
	for _, curso := range cursos {
		suma += curso.Nota
	}
	return suma / float64(len(cursos))
}

func ObtenerCursos(estudiantes []Estudiante) []string {
	cursosMap := make(map[string]bool)
	for _, estudiante := range estudiantes {
		for _, curso := range estudiante.Cursos {
			cursosMap[curso.Curso] = true
		}
	}
	cursos := make([]string, 0, len(cursosMap))
	for curso := range cursosMap {
		cursos = append(cursos, curso)
	}
	return cursos
}

func StringNoEstaEnSlice(target string, slice []string) bool {
	for _, s := range slice {
		if s == target {
			return false // El string está presente en el slice
		}
	}
	return true // El string no está presente en el slice
}
