package modelos

type Hombre struct {
	Edad       int     //publicas por mayuscula
	altura     float32 //privadas por minuscula
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	vivo       bool
}

func (h *Hombre) Respirar() {
	h.respirando = true
	h.vivo = true
}

func (h *Hombre) Pensar() {
	h.pensando = true
}

func (h *Hombre) Comer() {
	h.comiendo = true
}

func (h *Hombre) Sexo() string {
	return "Hombre"
}

func (h *Hombre) EsCarnivoro() bool {
	return true
}

func (h *Hombre) EstaVivo() bool {
	return h.vivo
}
