package variables

import "fmt"

func MostrarEnteros() {
	var byte byte = 255
	var entero int = 255
	var entero8 int8 = 127
	var entero16 int16 = 32767
	var entero32 int32 = 2147483647
	var entero64 int64 = 9223372036854775807

	fmt.Println("Byte:", byte)
	fmt.Println("Entero:", entero)
	fmt.Println("Entero 8 bits:", entero8)
	fmt.Println("Entero 16 bits:", entero16)
	fmt.Println("Entero 32 bits:", entero32)
	fmt.Println("Entero 64 bits:", entero64)
}

func AsignacionDirecta() {
	//El operador := es para declarar e inicializar una variable al mismo tiempo
	entero := 9223372036854775807
	entero8 := int8(127)
	entero16 := int16(32767)
	entero32 := int32(2147483647)
	entero64 := int64(9223372036854775807)

	fmt.Println("Entero:", entero)
	fmt.Println("Entero 8 bits:", entero8)
	fmt.Println("Entero 16 bits:", entero16)
	fmt.Println("Entero 32 bits:", entero32)
	fmt.Println("Entero 64 bits:", entero64)
}
