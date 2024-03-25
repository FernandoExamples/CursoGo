package modelos

type Mujer struct {
	Hombre
}

func (m *Mujer) Respirar()      { m.Respirando = true }
func (m *Mujer) Pensar()        { m.Pensando = true }
func (m *Mujer) Comer()         { m.Comiendo = true }
func (m *Mujer) Sexo() string   { return "Mujer" }
func (m *Mujer) EstaVivo() bool { return m.Vivo }
