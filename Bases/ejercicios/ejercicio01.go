package ejercicios

import (
	"strconv"
)

func Ejecrcicio1(value string) (int, string) {

	numero, _ := strconv.Atoi(value)

	if numero > 100 {
		return numero, "El numero es mayor a 100"
	}

	return numero, "El numero es menor a 100"
}
