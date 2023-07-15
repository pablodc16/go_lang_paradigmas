package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	saludo := "Hola mundo!"

	// Libreria strings
	fmt.Println("El mensaje original: ", saludo)
	fmt.Println("Existe 'Hola' dentro del string: ", strings.Contains(saludo, "Hola"))
	fmt.Println("Se reemplaza 'mundo' por 'amigos': ", strings.ReplaceAll(saludo, "mundo", "amigos"))
	fmt.Println("Imprime todo en mayuscula: ", strings.ToUpper(saludo))
	fmt.Println("Posicion de la letra 'm' en el string: ", strings.Index(saludo, "m"))
	fmt.Println("Dividir los elementos del string en un array: ", strings.Split(saludo, " "))

	// Libreria sort
	edades := []int{45, 20, 35, 30, 75, 60, 50, 25}

	fmt.Println("Slice original: ", edades)
	sort.Ints(edades) // este metodo modifica el slice original
	fmt.Println("Slice ordenado: ", edades)
	fmt.Println("Buscar la posicion del elemento '30': ", sort.SearchInts(edades, 30))

	nombres := []string{"Keren", "Sergio", "Hubal", "Pablo"}

	fmt.Println("Slice original: ", nombres)
	sort.Strings(nombres) // este metodo modifica el slice original
	fmt.Println("Slice ordenado: ", nombres)
	fmt.Println("Buscar la posicion del elemento 'Hubal': ", sort.SearchStrings(nombres, "Hubal"))
}

// para conocer mas sobre todas las librerias de go, por favor visitar el siguiente link: https://pkg.go.dev/std#stdlib
