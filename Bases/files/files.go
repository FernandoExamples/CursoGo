package files

import (
	"Bases/ejercicios"
	// "bufio"
	"fmt"
	"os"
)

func GuardarTabla() {
	var texto = ejercicios.Ejercicio2()
	archivo, err := os.Create("./files/txt/tabla.txt")

	if err != nil {
		println("Hubo un error al crear el archivo " + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumarTabla() {
	var texto = ejercicios.Ejercicio2()
	filename := "./files/txt/tabla.txt"
	if !Append(filename, texto) {
		fmt.Println("Error en la creacion del archivo")
	}

}

func Append(filename string, texto string) bool {
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return false
	}

	_, err = arch.WriteString(texto + "\n")
	if err != nil {
		fmt.Println(err)
		arch.Close()
		return false
	}

	arch.Close()

	return true
}

func LeerArchivo() {
	archivo, err := os.ReadFile("./files/txt/tabla.txt")
	if err != nil {
		fmt.Println("Hubo un error al leer el archivo " + err.Error())
		return
	}
	fmt.Println(string(archivo))
}
