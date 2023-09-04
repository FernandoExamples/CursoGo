package defer_ejemplo

import (
	"fmt"
	"log"
)

func ShowDefer() {
	fmt.Println("Este es un primer mensaje")

	defer fmt.Println("Este es un mensaje diferido")

	fmt.Println("Este es un segundo mensaje")
}

func EjemploPanic() {
	defer func ()  {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error que genero panic \n %v", reco)	
		}
	}()

	a := 1
	if a == 1 {
		panic("Se encontro el valor 1")
	}
}
