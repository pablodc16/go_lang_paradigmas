package main

import "fmt"

var puntaje = 105

func main() {
	// si pasamos la variable puntaje dentro de la funcion main se sale del scope del package main y entra en el de main, quedando inaccesible para las funciones del archivo saludos.go
	//var puntaje = 105

	saludar("Keren")

	for _, v := range puntos {
		fmt.Println(v)
	}

	mostrarPuntaje()
}
