package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Ejercicio2() string {
	scanner := bufio.NewScanner(os.Stdin)

	var numero int
	var err error
	var value string
	var texto string

	fmt.Println("Ingrese numero:")

	scanner.Scan()
	value = scanner.Text()
	numero, err = strconv.Atoi(value)

	for err != nil {
		fmt.Println("Ingrese numero:")
		scanner.Scan()
		value = scanner.Text()
		numero, err = strconv.Atoi(value)
	}

	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintln(numero, "x", i, "=", numero*i)
	}

	return texto
}
