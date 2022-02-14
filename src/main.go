package main

import (
	"BasicGolang/src/pc"
	"fmt"
)

func main() {
	a := 50
	b := &a
	a = 25

	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	myPc := pc.Pc{Ram: 16, Disk: 200, Brand: "msi"}
	fmt.Println(myPc)

	myPc.Ping()

	fmt.Println(myPc)
	myPc.DuplicateRam()

	fmt.Println(myPc)
	myPc.DuplicateRam()

	fmt.Println(myPc)

}
