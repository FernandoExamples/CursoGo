package variables

import (
	"fmt"
	"strconv"
	"time"
)

//Estas variables tienen escope en todo el paquete, incluso otros archivos porque comienzan en mayuscula

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Juan"
	Estado = true
	Sueldo = 500.50
	Fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func ConvertirATexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)

	return true, texto
}
