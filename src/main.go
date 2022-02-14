package main

import (
	"BasicGolang/src/pc"
	"fmt"
)

func main() {
	myPc := pc.Pc{Ram: 16, Brand: "MSI", Disk: 100}
	fmt.Print(myPc)
}
