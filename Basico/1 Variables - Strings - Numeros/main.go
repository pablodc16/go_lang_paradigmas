package main

import "fmt"

func main() {
	// *** strings
	var nombre1 string = "Keren"
	var nombre2 = "Jubal"
	var nombre3 string
	nombre4 := "Pablo" // este formato no se puede usar fuera de una funcion

	fmt.Println(nombre1, nombre2, nombre3, nombre4)

	nombre3 = "Sergio"

	fmt.Println(nombre1, nombre2, nombre3, nombre4)

	// *** ints
	var edad1 int = 20
	var edad2 = 30
	var edad3 int
	edad4 := 50

	fmt.Println(edad1, edad2, edad3, edad4)

	edad3 = 40

	fmt.Println(edad1, edad2, edad3, edad4)

	// Podemos utilizar variaciones del tipo INT para especificar el numero de bit que tendra de tamano en la memoria
	var num1 int8 = 25  // int 8 permite numeros desde -128 hasta 127
	var num2 uint = 5   // u int solo permite numeros positivos
	var num3 uint8 = 10 // u int 8 solo permite numeros positivos de 0 a 255

	fmt.Println(num1, num2, num3)

	// para saber mas sobre los tipos numericos, por favor utilice este link: https://go.dev/ref/spec#Numeric_types

	// *** float
	// para los tipos flotantes siempre hay que especificar la variacion y solo hay dos: 32 y 64
	var precio1 float32 = 100.15
	var precio2 float64 = 100000000.15
	var precio3 float32
	precio4 := 2.78

	fmt.Println(precio1, precio2, precio3, precio4)
}
