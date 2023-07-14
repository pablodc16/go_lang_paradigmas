package main

import "fmt"

func main() {
	edad := 45

	fmt.Println(edad <= 50)
	fmt.Println(edad >= 50)
	fmt.Println(edad == 45)
	fmt.Println(edad != 50)

	if edad < 30 {
		fmt.Println("La persona es menor de 30 años")
	} else if edad < 40 {
		fmt.Println("La persona es menor de 40 años")
	} else {
		fmt.Println("La persona tiene 45 años 0 mas")
	}
}
