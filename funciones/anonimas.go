package funciones

func Calculos() {

	var numero3 int = 10

	suma := func(numero1 int, numero2 int) int {
		return numero1 + numero2
	}

	println(suma(5, 5))

	suma = func(numero1 int, numero2 int) int {
		return numero1 + numero2 + numero3

	}

	println(suma(5, 5))

}
