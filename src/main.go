package main

import "fmt"

func main() {

	//Variables declaration
	helloMessage := "Hello"
	wordMessage := "world"

	// PrintLn
	fmt.Println(helloMessage, wordMessage)
	fmt.Println(helloMessage, wordMessage)

	// Printf
	nombre := "platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	// Sprintf
	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(message)

	// Type of data
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T", cursos)
}
