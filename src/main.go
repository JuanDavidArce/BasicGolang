package main

import (
	"BasicGolang/src/mypackage"
	"fmt"
)

func main() {
	var car mypackage.PublicCar
	car.Brand = "Toyota"
	car.Year = 2020
	fmt.Println(car)
}
