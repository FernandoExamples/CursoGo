package gorutines

import (
	"fmt"
	"strings"
	"time"
)

func MiNombreLento(nombre string, canal1 chan bool) {
	letras := strings.Split(nombre, "")

	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
	canal1 <- true
}
