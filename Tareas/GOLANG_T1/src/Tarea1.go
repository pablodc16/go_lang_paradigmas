package main

import (
	"fmt"
	"time"
)

func calcularEdad(añoNacimiento, mesNacimiento, díaNacimiento int) (años, meses, días, horas, minutos, segundos int) {
	fechaNacimiento := time.Date(añoNacimiento, time.Month(mesNacimiento), díaNacimiento, 0, 0, 0, 0, time.UTC)
	fechaActual := time.Now().UTC()

	// Calcular diferencia de tiempo
	diferencia := fechaActual.Sub(fechaNacimiento)

	// Calcular años completos
	años = int(diferencia.Hours() / 24 / 365.25)

	// Calcular meses completos
	fechaCumpleañosActual := fechaNacimiento.AddDate(años, 0, 0)
	for fechaCumpleañosActual.Before(fechaActual) {
		fechaCumpleañosActual = fechaCumpleañosActual.AddDate(0, 1, 0)
		meses++
	}

	fechaCumpleañosActual = fechaCumpleañosActual.AddDate(0, -1, 0)
	días = int(fechaActual.Sub(fechaCumpleañosActual).Hours() / 24)

	// Calcular horas, minutos y segundos
	horas = int(diferencia.Hours()) % 24
	minutos = int(diferencia.Minutes()) % 60
	segundos = int(diferencia.Seconds()) % 60

	return años, meses, días, horas, minutos, segundos
}

func main() {
	var año, mes, día int
	fmt.Print("Ingresa el año de nacimiento: ")
	fmt.Scanln(&año)
	fmt.Print("Ingresa el mes de nacimiento (1-12): ")
	fmt.Scanln(&mes)
	fmt.Print("Ingresa el día de nacimiento: ")
	fmt.Scanln(&día)

	años, meses, días, horas, minutos, segundos := calcularEdad(año, mes, día)

	fmt.Println("Tu edad es:")
	fmt.Println("Años:", años)
	fmt.Println("Meses:", meses)
	fmt.Println("Días:", días)
	fmt.Println("Horas:", horas)
	fmt.Println("Minutos:", minutos)
	fmt.Println("Segundos:", segundos)
}
