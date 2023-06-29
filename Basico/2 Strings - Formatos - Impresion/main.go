package main

import "fmt"

func main() {
	// Print (no agrega salto de linea)
	fmt.Print("Hola ")
	fmt.Print("Mundo \n") // el salto de linea se agrega con el \n
	fmt.Print("siguiente linea \n")

	// Println (agrega salto de linea automaticamente al final de string)
	fmt.Println("Hola Paradigmas")
	fmt.Println("Adios Paradigmas")

	// Concatenar en el Print
	nombre := "Pablo"
	edad := 36
	fmt.Println("Hola mi nombre es", nombre, "y tengo", edad, "años de edad")

	// Printf (formatos de String)
	// se utiliza un format specifier para darle formato a las variables que debe sustituir

	fmt.Printf("Hola mi nombre es %v y tengo %v años de edad \n", nombre, edad) // hace una simple sustitucion

	fmt.Printf("Hola mi nombre es %q y tengo %q años de edad \n", nombre, edad) // ademas de sustituir, envuelve las variables entre comillas, las variables deben ser string, no funciona con otro formato

	fmt.Printf("La edad es de tipo %T \n", edad) // imprime el tipo de dato que es la variable

	fmt.Printf("El precio del producto es %f \n", 105.55) // imprime la variable como un float

	fmt.Printf("El precio del producto es %0.2f \n", 105.557) // imprime la variable como un float y limitamos el numero de decimales a dos (hace un redondeo)

	// para conocer todos los format specifiers, por favor use el siguiente link: https://pkg.go.dev/fmt

	// Sprintf (salva los string formateados en una variable)
	var mensaje = fmt.Sprintf("Hola mi nombre es %v y tengo %v años de edad \n", nombre, edad)
	fmt.Println("El mensaje salvado es: ", mensaje)
}
