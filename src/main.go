package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20
	fmt.Println(m)

	// Recorrer un map
	for indice, valor := range m {
		fmt.Println(indice, valor)
	}

	// Encontrar un valor
	valor, ok := m["Jose"]
	fmt.Println(valor, ok)
}
