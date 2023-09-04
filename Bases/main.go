package main

import (
	"Bases/ejercicios"
	"Bases/mapas"
	"Bases/teclado"
	"Bases/variables"
	"fmt"
	"runtime"
)

func Bases() {
	fmt.Println("Hello, World!")

	fmt.Println("----------------------- Enteros -----------------------")
	variables.MostrarEnteros()

	fmt.Println("----------------------- Resto -----------------------")
	variables.RestoVariables()

	fmt.Println("----------------------- Conversiones -----------------------")
	fmt.Println(variables.ConvertirATexto(65))

	fmt.Println("----------------------- Condicionales -----------------------")

	os := runtime.GOOS

	//Esto es un if normal
	if os == "windows" {
		fmt.Println("Estas en windows")
	}

	//Esto es un if con inicializacion de variable
	if os2 := runtime.GOOS; os2 == "windows" {
		fmt.Println("Estas en windows")
	}

	//Esto es un switch
	switch os {
	case "windows":
		fmt.Println("Estas en windows")
	case "darwin":
		fmt.Println("Estas en mac")
	default:
		fmt.Println("No se encontro el sistema operativo")
	}

	//Esto es un switch con inicializacion de variable
	switch os2 := runtime.GOOS; os2 {
	case "windows":
		fmt.Println("Estas en windows")
	case "darwin":
		fmt.Println("Estas en mac")
	default:
		fmt.Println("No se encontro el sistema operativo")
	}

	fmt.Println("----------------------- Bucles -----------------------")

	//Esto es un for normal
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//Esto es un while
	i := 4
	for i >= 0 {
		i--
		fmt.Println(i)
	}

	//bucle for range (foreach)
	slice := []int{1, 2, 3, 4, 5}
	for index, value := range slice {
		fmt.Println(index, value)
	}

	fmt.Println("----------------------- Ejercicio 1 -----------------------")
	numero, texto := ejercicios.Ejercicio1("101")
	fmt.Println(numero, texto)
}

func Teclado() {
	teclado.IngresoNumeros()
}

func main() {
	// Bases()
	// Teclado()
	// println(ejercicios.Ejercicio2())
	// files.GuardarTabla()
	// files.SumarTabla()
	// files.LeerArchivo()
	// funciones.Calculo()
	// funciones.LlamarClosure()
	// println(funciones.Fibonacci(10))
	// arreglos.MostrarArreglos()
	// arreglos.MostrarSlices()
	// arreglos.Capacidad()
	mapas.MostrarMapas()
}
