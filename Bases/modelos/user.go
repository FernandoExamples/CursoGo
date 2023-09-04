package modelos

import (
	"time"
)

type Usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

func (usuario *Usuario) SetData(id int, nombre string, fechaalta time.Time, status bool) {
	usuario.Id = id
	usuario.Nombre = nombre
	usuario.FechaAlta = fechaalta
	usuario.Status = status
}
