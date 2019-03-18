package main

import (
	"fmt"
	"gosol/templates"
	"os"
)

func help() {
	fmt.Println("./gosol comp  #complier solidity contract")
	fmt.Println("./gosol build  #build test code")
}

func main() {
	if len(os.Args) < 2 {
		help()
		os.Exit(0)
	}
	if os.Args[1] == "comp" {
		ParseRun()
	} else if os.Args[1] == "build" {
		//
		templates.BuildRun()
	} else {
		help()
		os.Exit(0)
	}
}
