package mapas

import "fmt"

func MostrarMapas() {
	paises := map[string]string{"Mexico": "D.F.", "Argentina": "Buenos Aires", "Colombia": "Bogota", "Peru": "Lima"}
	paises["Chile"] = "Santiago de Chile"

	fmt.Println(paises)
	fmt.Println(paises["Mexico"])

	campeonatos := map[string]int{"Barcelona": 39, "Real Madrid": 33, "Chivas": 12, "Boca Juniors": 26}
	delete(campeonatos, "Real Madrid")
	
	for equipo, puntaje := range campeonatos {
		fmt.Println(equipo, puntaje)
	}

	puntaje, existe := campeonatos["Mineiro"]
	fmt.Println(puntaje, existe)
}
