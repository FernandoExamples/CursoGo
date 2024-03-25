package ejercicios

import "strconv"

func Ejercicio1(numero string) (int, string) {
	valor, err := strconv.Atoi(numero)

	if err != nil {
		return -1, err.Error()
	}

	if valor > 100 {
		return 0, "Es mayor a 100"
	} else {
		return 0, "Es menor a 100"
	}
}
