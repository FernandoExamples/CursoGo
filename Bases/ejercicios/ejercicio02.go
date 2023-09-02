package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Ejercicio2() {
	scanner := bufio.NewScanner(os.Stdin)

	var numero int
	var err error
	var value string

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
		fmt.Println(value, "*", i, "=", numero*i)
	}
}

