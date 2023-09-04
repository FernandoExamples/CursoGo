package modelos

type Mujer struct {
	Hombre
}

func (m *Mujer) Sexo() string {
	return "Mujer"
}
