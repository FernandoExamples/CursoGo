package variables

import (
	"fmt"
	"time"
	"strconv"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func MostarCadenas() {
	var cadena string = "Hola, Mundo!"
	var caracter byte = 'A'
	var cadenaMultilinea string = `Hola, 
	Mundo!`

	fmt.Println("cadena:", cadena)
	fmt.Println("caracter:", caracter)
	fmt.Println("cadenaMultilinea:", cadenaMultilinea)
}

func MostrarGlobales() {
	Nombre = "Juan"
	Estado = true
	Sueldo = 1000.50
	Fecha = time.Now()

	fmt.Println("Nombre:", Nombre)
	fmt.Println("Estado:", Estado)
	fmt.Println("Sueldo:", Sueldo)
	fmt.Println("Fecha:", Fecha)
}

func ToText(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}
