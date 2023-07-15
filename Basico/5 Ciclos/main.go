package main

import "fmt"

func main() {
	x := 0

	for x < 5 {
		fmt.Println("El valor de x es: ", x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("El valor de i es: ", i)
	}

	nombres := []string{"Keren", "Sergio", "Hubal", "Pablo"}

	for index, value := range nombres {
		fmt.Printf("El valor en la posicion %v es %v \n", index, value)
	}

	// en caso de que no quisieramos usar index o value, se pone un _
	for _, value := range nombres {
		fmt.Printf("El valor es %v \n", value)
	}

	// continue y break
	for index, value := range nombres {

		if index == 1 {
			// si usamos continue, se deja de ejecutar todo el codigo que esta por debajo y continua con el siguiente indice
			fmt.Println("Nos brincamos el indice 1")
			continue
		}

		if index > 2 {
			// si usamos break, se detiene completamente el ciclo
			fmt.Println("Se acaba el ciclo aunque hayan mas elementos en el slice.")
			break
		}

		fmt.Printf("El valor en la posicion %v es %v \n", index, value)
	}
}
