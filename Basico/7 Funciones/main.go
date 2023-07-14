package main

import (
	"fmt"
	"math"
	"strings"
)

func saludar(n string) {
	fmt.Printf("Buenos dias %v \n", n)
}

func despedir(n string) {
	fmt.Printf("Adios %v \n", n)
}

func cicloNombres(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

// funcion para retornar 1 valor
func areaCirculo(r float64) float64 {
	return math.Pi * r * r
}

// funcion para retornar 2 valores
func obtenerIniciales(n string) (string, string) {
	s := strings.ToUpper(n)
	nombres := strings.Split(s, " ")

	var iniciales []string

	for _, v := range nombres {
		iniciales = append(iniciales, v[:1]) // v[:1] me devuelve la primera letra del valor que esta en v
	}

	if len(iniciales) > 1 {
		return iniciales[0], iniciales[1]
	}

	return iniciales[0], "_"
}

func main() {
	saludar("Hubal")
	despedir("Keren")

	nombres := []string{"Keren", "Sergio", "Hubal", "Pablo"}
	cicloNombres(nombres, saludar)

	a1 := areaCirculo(10.5)
	a2 := areaCirculo(15)

	fmt.Printf("El area del circulo 1 es %0.3f y el area del circulo 2 es %0.3f \n", a1, a2)

	// como la funcion retorna 2 valores, necesitamos 2 variables para almacenarlos
	pn1, sn1 := obtenerIniciales("pablo duran")
	pn2, sn2 := obtenerIniciales("Keren")
	fmt.Println("Las iniciales del nombre 1 son: ", pn1, sn1)
	fmt.Println("Las iniciales del nombre 2 son: ", pn2, sn2)
}
