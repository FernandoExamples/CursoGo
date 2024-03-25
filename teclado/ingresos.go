package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumeros() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese el primer número: ")

	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado no es un número")
		}
	}

	fmt.Println("Ingrese el segundo número: ")

	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado no es un número")
		}
	}

	fmt.Println("Ingrese la leyenda: ")

	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	fmt.Printf("Los números ingresados son: %d y %d\n", numero1, numero2)
	fmt.Printf("La leyenda ingresada es: %s\n", leyenda)

}
