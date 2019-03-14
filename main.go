package main

import (
	"fmt"
	"gosol/toml"
)

func main() {
	fmt.Println("hello world")
	//ParseJson("pdmall.abi")

	solfiles, err := ParseDir("sol")
	fmt.Println(solfiles, err)
	for _, solfile := range solfiles {
		fmt.Println(solfile)
		//Compiler(solfile, "sol", "contracts")
	}
	fmt.Printf("%+v\n", toml.Config.Eth)
}
