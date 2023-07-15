package main

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
