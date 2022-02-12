package main

import "fmt"

func main() {

	// Area Cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("area cuadrado: ", areaCuadrado)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("suma ", result)

	// Rexta
	result = y - x
	fmt.Println("Resta", result)

	// Multiplicacion
	result = x * y
	fmt.Println("multiplicacion:", result)

	// Division
	result = y / x
	fmt.Println("division: ", result)

	// Modulo
	result = y % x
	fmt.Println("Modulo: ", result)

	// Incremental
	x++
	fmt.Println("Incremental:", x)

	// Decremental
	x--
	fmt.Println("Decremental: ", x)
}
