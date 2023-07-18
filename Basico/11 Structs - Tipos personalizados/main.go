package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func obtenerInput(promt string, r *bufio.Reader) (string, error) {
	fmt.Print(promt)

	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func crearFactura() factura {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Digite el nombre de la factura: ")
	// nombre, _ := reader.ReadString('\n')
	// nombre = strings.TrimSpace(nombre)

	nombre, _ := obtenerInput("Digite el nombre de la factura: ", reader)

	f := nuevaFactura(nombre)

	fmt.Println("Se creo la factura para ", f.nombre)

	return f
}

func promtOptions(f factura) {
	reader := bufio.NewReader(os.Stdin)

	opc, _ := obtenerInput("Elija una opcion (a - Agregar item, g - Guardar factura, t - Agregar tip): ", reader)

	switch opc {
	case "a":
		nombre, _ := obtenerInput("Nombre del item: ", reader)
		precio, _ := obtenerInput("Precio del item: ", reader)

		precioParseado, err := strconv.ParseFloat(precio, 64)

		if err != nil {
			fmt.Println("El precio debe ser un numero")
			promtOptions(f)
		}

		f.agregarItem(nombre, precioParseado)

		fmt.Println("Item agregado - ", nombre, precio)
		promtOptions(f)
	case "t":
		tip, _ := obtenerInput("Digite el monto del tip: ", reader)

		tipParseado, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("El tip debe ser un numero")
			promtOptions(f)
		}

		f.actualizarTip(tipParseado)

		fmt.Println("Tip agregado - ", tip)
		promtOptions(f)
	case "g":
		f.salvarFactura()

		fmt.Println("Salvaste la factura - ", f.nombre)
	default:
		fmt.Println("Opcion invalida")
		promtOptions(f)
	}
}

func main() {
	miFactura := crearFactura()
	promtOptions(miFactura)

	// fmt.Println(miFactura)

	// miFactura.agregarItem("Sopa", 4.50)
	// miFactura.agregarItem("Pie", 8.95)
	// miFactura.agregarItem("Ensalada", 4.95)
	// miFactura.agregarItem("Cafe", 3.25)

	// miFactura.actualizarTip(10)

	// fmt.Println(miFactura.formatoFactura())

}
