package arreglos

import "fmt"

func MostrarArreglos() {
	var tabla = [10]int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	tabla[0] = 1
	fmt.Println(tabla)

	tabla2 := [10]string{"colombia", "peru", "argentina"}

	for i := 0; i < len(tabla2); i++ {
		fmt.Println(tabla2[i])
	}

	var matriz [5][7]int

	matriz[3][5] = 55
	fmt.Println(matriz)
}
