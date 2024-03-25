package arrreglos

import "fmt"

var tablaS []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{5, 6, 7, 8, 9, 10}

func MostrarSlices() {
	fmt.Println(tablaS)

	porcion := arreglo[3:] // desde la posicion 3 hasta el final

	porcion2 := arreglo[:5] // desde el inicio hasta la posicion 5

	porcion3 := arreglo[3:7] // desde la posicion 3 hasta la 7

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)

}

func Capacidad() {
	elementos := make([]int, 5, 20)

	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("\nLargo %d, Capacidad %d\n", len(nums), cap(nums))
}
