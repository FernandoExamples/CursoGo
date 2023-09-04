package ejercicios

import (
	"Bases/interfaces"
	"fmt"
)

func Interfaces(humano interfaces.Humano) {

	humano.Respirar()

	fmt.Printf("Soy un/a %s y estoy respirando \n", humano.Sexo())

}
