package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
	//ParseJson("pdmall.abi")

	solfiles, err := ParseDir("sol")
	fmt.Println(solfiles, err)
	for _, solfile := range solfiles {
		Compiler(solfile, "sol", "contracts")
	}
}
