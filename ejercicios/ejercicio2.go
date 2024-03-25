package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaMultiplicar() string {
	var numero int
	var texto string

	scanner := bufio.NewScanner(os.Stdin)

	for {
		println("Ingrese un número: ")
		if scanner.Scan() {
			var err error
			numero, err = strconv.Atoi(scanner.Text())

			if err == nil {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		resultado := numero * i
		texto += fmt.Sprintf("%d x %d = %d\n", numero, i, resultado)
	}

	return texto
}
