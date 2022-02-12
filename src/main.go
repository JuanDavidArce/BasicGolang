package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)

}

func triplaArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c int, d int) {
	return a, a * 2
}

func main() {

	normalFunction("hola mundo")
	triplaArgument(1, 2, "hola")
	value := returnValue(2)
	fmt.Println("value: ", value)

	value1, value2 := doubleReturn(2)
	fmt.Println("Value1 ", value1, "Value2 ", value2)
}
