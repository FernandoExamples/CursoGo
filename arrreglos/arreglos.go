package arrreglos

import "fmt"

var tabla [10]int = [10]int{5, 6, 7, 8, 9, 10}
var matriz [5][7]int

func MostrarArreglos() {
	tabla[0] = 1
	tabla[5] = 15

	tabla2 := [10]string{"uno", "dos", "tres", "cuatro", "cinco"}

	fmt.Println(tabla)
	fmt.Println(tabla2)

	for _, valor := range tabla {
		fmt.Println(valor)
	}

	for i := 0; i < len(tabla2); i++ {
		fmt.Println(tabla2[i])
	}

	matriz[2][5] = 1
	matriz[3][1] = 1

	fmt.Println(matriz)
}
