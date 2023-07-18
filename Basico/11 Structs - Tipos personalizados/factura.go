package main

import (
	"fmt"
	"os"
)

// Un struct (abreviado de la palabra structure) sirve para crear un miembro con diferentes tipos de datos dentro de una misma variable, a diferencia de los Array que solo almacenan datos del mismo tipo. No es una clase (GO no usa clases)

type factura struct {
	nombre string
	items  map[string]float64
	tip    float64
}

func nuevaFactura(nombre string) factura {
	fact := factura{
		nombre: nombre,
		items:  map[string]float64{},
		tip:    0,
	}

	return fact
}

// funcion para agregar datos a la factura
// utilizamos un "Receiver" para que esta funcion este siempre asociada con los struct facturas y no sea llamado por cualquier otro tipo
func (f *factura) formatoFactura() string {
	// *** Esta funcion tambien funciona sin el puntero en el "Receiver", pero podria afectar el rendimiento de la aplicacion si se usa constantemente, ya que crearia una copia cada vez que se llama, lledando la memoria de muchas de ellas

	ff := "Detalle de la factura: \n"
	var total float64 = 0

	// agregar los items a la variable ff
	for k, v := range f.items {
		ff += fmt.Sprintf("%-25v ...$%v \n", k+":", v) // con el '-25' despues del signo de '%' le decimos que tome un espacio de 25 caracteres, o sea, la palabra que es el key (cafe, pie, etc) mas lo que haga falta para llegar a 25 en espacios en blanco, es una especie de tabulado, asi los precios se van a ver alineados
		total += v
	}

	// agregar el Tip a la factura
	ff += fmt.Sprintf("%-25v ...$%0.2f \n", "Tip:", f.tip)

	// agregar el total a la factura
	ff += fmt.Sprintf("%-25v ...$%0.2f \n", "Total sin Tip:", total)

	// total + tip
	ff += fmt.Sprintf("%-25v ...$%0.2f \n", "Total con Tip:", total+f.tip)

	return ff
}

// como se vio en el modulo de "Pass by value", los struct son elementos que no se actualizan cuando se pasan por parametro, lo que se pasa es una copia de ese elemento, es por eso que debemos pasar un puntero con la direccion del elemento para que estas funciones de abajo actualizen el elemento como tal, esto se logra agregando un * al tipo en el "Receiver"

func (f *factura) actualizarTip(tip float64) {
	f.tip = tip
}

func (f *factura) agregarItem(nombre string, precio float64) {
	f.items[nombre] = precio
}

func (f *factura) salvarFactura() {
	datos := []byte(f.formatoFactura())

	error := os.WriteFile("facturas/"+f.nombre+".txt", datos, 0644)

	if error != nil {
		panic(error)
	}

	fmt.Println("La factura se salvo.")
}
