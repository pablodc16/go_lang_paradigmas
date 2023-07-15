package main

import "fmt"

// una mapa es un objeto que consta de un Key y un Value

func main() {
	menu := map[string]float64{
		"Sopa":     4.99,
		"Pie":      7.99,
		"Ensalada": 6.99,
		"Cafe":     2.50,
	}

	fmt.Println(menu)
	fmt.Println(menu["Pie"])

	// actualizar un Value, *no se actualizan los Key
	menu["Pie"] = 5.99

	// ciclo dentro de un mapa
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}
}

// Para saber mas a detalle como funcionan los mapas, por favor visite el siguiente link: https://go.dev/blog/maps
