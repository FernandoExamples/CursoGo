package ejercicios

import (
	"cursogo/interfaces"
	"fmt"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()

	fmt.Printf("Soy un/a %s y respiro\n", hu.Sexo())
}
