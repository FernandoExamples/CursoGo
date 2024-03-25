package funciones

func tabla(valor int) func() int {
	secuencia := 0

	return func() int {
		secuencia++
		return valor * secuencia
	}
}

func LlamarClosure() {
	tablaDel := 2

	miTabla := tabla(tablaDel)

	for i := 0; i < 10; i++ {
		println(miTabla())
	}
}
