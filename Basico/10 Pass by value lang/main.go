package main

import "fmt"

func actualizarNombre(x string) {
	x = "Sergio"
}

func actualizarMenu(y map[string]float64) {
	y["Cafe"] = 3.51
	y["Hamburguesa"] = 8.25
}

func main() {
	// Grupo A de tipos -> string, int, bool, float, array, struct
	nombre := "Hubal"

	actualizarNombre(nombre) // aqui le estamos pasando una copia de la variable a la funcion, no la variable como tal

	fmt.Println(nombre) // al imprimir la variable 'nombre' sigue mostrando el valor 'Hubal' a pesar de que se intento cambiarlo en la funcion actualizarNombre, la razon de esto es lo que se explica en el comentario anterior

	// Grupo B de tipos -> slice, map, function
	menu := map[string]float64{
		"Sopa":     4.99,
		"Pie":      7.99,
		"Ensalada": 6.99,
		"Cafe":     2.50,
	}

	actualizarMenu(menu) // en este caso lo que se envia a la funcion es un puntero con la direccion en memoria de donde se encuentra este tipo (el map), es por ello que en este caso si se actualiza el valor, e incluso agregamos uno nuevo al map llamado menu
	fmt.Println(menu)
}
