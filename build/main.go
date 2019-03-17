package main

import (
	"os"
)

func main() {
	if len(os.Args) < 2 {
		Usage()
		os.Exit(0)
	}
	if os.Args[1] == "1" {
		ImplRunCode()
	} else {
		Run()
	}

}
