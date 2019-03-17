
package main

import (
	"fmt"
	"os"
)

func Usage() {
	fmt.Printf("1 - build   :%s  1\n", os.Args[0])
	fmt.Printf("2 - deploy  :%s  2\n", os.Args[0])
	num := 3
	for _, v := range Config.FuncConfs {
		fmt.Printf("%d - test %s:%s  %d\n", num, v.Func, os.Args[0], num)
		num++
	}
	for _, v := range Config.NoGasFuncConfs {
		fmt.Printf("%d - test %s:%s  %d\n", num, v.Func, os.Args[0], num)
		num++
	}
}

func Run() {
	if os.Args[1] == "2" {
		DeployPdbank()
	} else if os.Args[1] == "3" {
		CallDeposit("0x70c53a4c94ccce9ce56effbfcb89b221f986cd41", "123")
	} else if os.Args[1] == "4" {
		CallWithdraw("0x70c53a4c94ccce9ce56effbfcb89b221f986cd41", "123")
	} else if os.Args[1] == "5" {
		CallBankName()
	} else if os.Args[1] == "6" {
		CallBalances()
	} 
}