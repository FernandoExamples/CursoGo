package users

import (
	"Bases/modelos"
	"fmt"
	"time"
)

func AltaUsuario() {
	u := new(modelos.Usuario)
	u.SetData(1, "Juan", time.Now(), true)
	fmt.Println(u)
}
