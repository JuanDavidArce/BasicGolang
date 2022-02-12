package main

import "fmt"

func main() {

	// Constants declaration
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("Pi", pi)
	fmt.Println("pi2", pi2)

	// Variables declaration
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero Values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area Cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("area cuadrado: ", areaCuadrado)
}
