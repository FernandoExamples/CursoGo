package archivos

import (
	"fmt"
	"os"
)

var fileName = "./texto.txt"

func GuardarTexto(text string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.WriteString(text)
}

func AppedTexto(text string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.WriteString(fmt.Sprintf("\n%s", text))
}

func LeerArchivo() {
	bs, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bs))
}
