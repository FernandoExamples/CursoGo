package funciones

func Calculo() {

	var numero3 int = 32
	var numero4 int = 14
	
	suma := func(numero1 int, numero2 int) int {
		return numero1 + numero2 + numero3 + numero4
	}

	println(suma(5, 5))
}
