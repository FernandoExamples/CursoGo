package variables

import "fmt"

func MostarEnteros() {
	var _entero int = 42
	var _int32 int32 = 42
	var _int64 int64 = 42
	var _uint32 uint32 = 42
	var _uint64 uint64 = 42

	fmt.Println("int =", _entero)
	fmt.Println("int de 32 =", _int32)
	fmt.Println("int de 64 =", _int64)
	fmt.Println("unsigned int de 32 =", _uint32)
	fmt.Println("unsigned int de 64 =", _uint64)
}
