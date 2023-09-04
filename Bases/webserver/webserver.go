package webserver

import (
	"fmt"
	"net/http"
)

func MiWebServer() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":3000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./webserver/index.html")
}

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func Middleware() {
	fmt.Println("Ejecutando Middleware")

	result := operacionesMidd(sumar)(5, 5)
	fmt.Println(result)

	result = operacionesMidd(restar)(5, 5)
	fmt.Println(result)

	result = operacionesMidd(multiplicar)(5, 5)
	fmt.Println(result)
}

func operacionesMidd(f func(a, b int) int) func(a, b int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operación")
		return f(a, b)
	}
}
