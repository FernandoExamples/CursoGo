package middlewares

import (
	"fmt"
)

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func MiMiddleware() {

	result := operacionesMid(sumar)(1, 2)
	fmt.Println(result)
	result = operacionesMid(restar)(1, 2)
	fmt.Println(result)
	result = operacionesMid(multiplicar)(1, 2)
	fmt.Println(result)

}

func operacionesMid(f func(int, int) int) func(int, int) int {
	return func(i1, i2 int) int {
		fmt.Println("Inicio operaciones")
		return f(i1, i2)
	}
}
