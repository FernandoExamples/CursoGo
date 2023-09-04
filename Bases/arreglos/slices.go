package arreglos

import "fmt"

func MostrarSlices() {
	var tabla = []int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	var slice = tabla[2:5]
	var slice2 = tabla[:5]
	var slice3 = tabla[2:]

	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(slice3)
}

func Capacidad() {
	tabla := make([]int, 5, 10)
	fmt.Printf("Largo %d, Capacidad %d", len(tabla), cap(tabla))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {

		nums = append(nums, i)
	}

	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))

}
