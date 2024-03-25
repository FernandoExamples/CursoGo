package gorutines

import (
	"strings"
	"time"
)

func MiNombreLento(nombre string, canal chan bool) {
	letras := strings.Split(nombre, "")

	for _, letra := range letras {
		time.Sleep(500 * time.Millisecond)
		println(letra)
	}

	canal <- true
}
