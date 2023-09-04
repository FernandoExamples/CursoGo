package funciones

func tabla(valor int) func() int {
	numero := valor
	secuencia := 0

	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}

func LlamarClosure() {
	tabla2 := tabla(2)

	for i := 0; i <= 10; i++ {
		println(tabla2())
	}

}
