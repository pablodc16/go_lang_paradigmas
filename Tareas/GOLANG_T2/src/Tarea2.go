package main

import "fmt"

func main() {
	var nombreEstudiante string
	notas := make([]float64, 5)

	fmt.Print("Ingresa el nombre del estudiante: ")
	fmt.Scanln(&nombreEstudiante)

	for i := 0; i < 5; i++ {
		var materia string
		var nota float64

		fmt.Printf("Ingresa el nombre de la materia #%d: ", i+1)
		fmt.Scanln(&materia)
		fmt.Printf("Ingresa la nota obtenida en %s: ", materia)
		fmt.Scanln(&nota)

		notas[i] = nota
	}

	promedio := calcularPromedio(notas)

	fmt.Printf("El promedio final de %s es: %.2f\n", nombreEstudiante, promedio)
}

func calcularPromedio(notas []float64) float64 {
	suma := 0.0
	for _, nota := range notas {
		suma += nota
	}
	promedio := suma / float64(len(notas))
	return promedio
}
