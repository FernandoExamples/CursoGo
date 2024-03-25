package defer_panic

import (
	"fmt"
	"log"
)

/**
	defer se utiliza para posponer una isntrucción hasta antes de una instrucción return sea ejecutada. Usualmente se usa para cerrar recursos de IO, como archivos o conexiones a bases de datos.
	Si hay varias instrucciones defer, se ejecutarán en orden inverso al que fueron declaradas, como una pila.
**/

func EjemploDefer() {
	defer fmt.Println("Hola")
	fmt.Println("Mundo")
}

func EjemploPanic() {

	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error: %v", reco)
		}
	}()

	a := 1
	if a == 1 {
		panic("Error de red")
	}
}
