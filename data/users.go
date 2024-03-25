package data

import (
	"cursogo/modelos"
	"fmt"
	"time"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.SetData(1, "jorge", "123", "jorge@gmail.com", time.Now(), true)
	fmt.Println(u)
}
