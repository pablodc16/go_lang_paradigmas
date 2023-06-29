package main

import "fmt"

func main() {
	// Array (tienen un tamaño fijo)
	var edades [3]int = [3]int{20, 25, 30}
	var precios = [3]int{205, 125, 320}
	nombres := [4]string{"Keren", "Jubal", "Sergio", "Pablo"}

	fmt.Println(edades)
	fmt.Println(edades, len(edades))
	fmt.Println(precios)
	fmt.Println(nombres)

	nombres[1] = "Michael"
	fmt.Println(nombres)

	// Slices (Se puede manipular su tamaño)
	var puntajes = []int{100, 150, 20}
	fmt.Println(puntajes)

	puntajes[2] = 50
	fmt.Println(puntajes)

	puntajes = append(puntajes, 85)
	fmt.Println(puntajes)

	// Slice Ranges (permite devolver los elementos dentro de un rango)
	rango1 := nombres[1:3] // devuelve los elementos del 1 al 3 en el array, incluye el 1 pero no incluye el 3
	rango2 := nombres[2:]  // devuelve los elementos del 2 hasta el ultimo en el array
	rango3 := nombres[:3]  // devuelve los elementos desde el primero hasta el 3, sin incluir el 3
	fmt.Println(rango1, rango2, rango3)
}
