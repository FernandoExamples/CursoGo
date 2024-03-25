package main

import (
	"cursogo/arrreglos"
	"cursogo/data"
	"cursogo/defer_panic"
	"cursogo/ejercicios"
	"cursogo/funciones"
	"cursogo/iteraciones"
	"cursogo/mapas"
	"cursogo/middlewares"
	"cursogo/modelos"
	"cursogo/variables"
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello, World!")
	variables.MostarEnteros()
	variables.MostrarReales()
	variables.MostarCadenas()
	variables.MostrarGlobales()
	fmt.Println(variables.ToText(1))

	if os := runtime.GOOS; os == "linux" {
		fmt.Println("Linux")
	} else {
		fmt.Println("No es Linux")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux")
	case "darwin":
		fmt.Println("MacOS")
	default:
		fmt.Printf("%s \n", os)
	}

	_, msg := ejercicios.Ejercicio1("100")

	fmt.Println(msg)

	iteraciones.For()

	// archivos.AppedTexto(ejercicios.TablaMultiplicar())
	// archivos.LeerArchivo()
	// teclado.IngresoNumeros()

	funciones.Calculos()
	funciones.LlamarClosure()

	fmt.Println("-------------------Recursividad-----------------")
	funciones.Exponencia(2)
	fmt.Println("-------------------Arreglos-----------------")
	arrreglos.MostrarArreglos()
	arrreglos.MostrarSlices()
	arrreglos.Capacidad()

	fmt.Println("-------------------Mapas-----------------")

	mapas.MostrarMapas()

	fmt.Println("-------------------POO-----------------")
	data.AltaUsuario()

	pedro := new(modelos.Hombre)
	ejercicios.HumanosRespirando(pedro)

	maria := new(modelos.Mujer)
	ejercicios.HumanosRespirando(maria)

	fmt.Println("-------------------DEFER-----------------")
	defer_panic.EjemploDefer()
	// defer_panic.EjemploPanic()

	fmt.Println("-------------------GoRutinas-----------------")

	// canal := make(chan bool)

	// go gorutines.MiNombreLento("Fernando", canal)

	// estado := <-canal

	// fmt.Println("La rutina terminó", estado)

	// webserver.Server()
	middlewares.MiMiddleware()
}
