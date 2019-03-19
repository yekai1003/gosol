
package main

import (
	"fmt"
	"os"
)

type FuncsList struct {
	FuncName string
	Num      int
}

func Usage() {
	fmt.Printf("1 - help   :%s  1\n", os.Args[0])
	fmt.Printf("2 - deploy  :%s  2\n", os.Args[0])
	num := 3
	for _, v := range datas {
		fmt.Printf("%d - test %s:%s  %d\n", num, v.FuncName, os.Args[0], num)
		num++
	}
}

func Run() {
	if len(os.Args) < 2 || os.Args[1] == "1" {
		Usage()
		os.Exit(0)
	}
	if os.Args[1] == "2" {
		DeployPdbank()
	} else if os.Args[1] == "3" {
		CallTotalAmount()
	} else if os.Args[1] == "4" {
		CallBankName()
	} else if os.Args[1] == "5" {
		CallBalances()
	} else if os.Args[1] == "6" {
		CallWithdraw(Config.Deploy.TestAddr, Config.Deploy.Pass)
	} else if os.Args[1] == "7" {
		CallOwner()
	} else if os.Args[1] == "8" {
		CallDeposit(Config.Deploy.TestAddr, Config.Deploy.Pass)
	} 
}
var datas = []FuncsList{
	{"totalAmount",3},
	{"bankName",4},
	{"balances",5},
	{"withdraw",6},
	{"owner",7},
	{"deposit",8},

}


func main() {
	Run()
}

