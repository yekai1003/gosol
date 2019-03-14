package main

import (
	"fmt"
	"os"
	"text/template"
)

type DeployContract struct {
	ContractCall string
	ContratName  string
}

const axx = `func DeployContract(addr, pass, keystorePath string) (common.Address, error) {
	auth, err := MakeAuth(addr, pass, keystorePath)
	if err != nil {
		fmt.Println("failed to makeAuth", err)
		return nil, err
	}
	//common.Address, *types.Transaction, *Pdbank, error
	contractaddr, _, _, err := contracts.{{.ContractCall}}
	if err != nil {
		fmt.Println("failed to %s",{{.ContratName}}, err)
		return nil, err
	}
	return contractaddr, err
}`

func main() {
	x1 := DeployContract{"(auth,cli)", "xxx.youcan"}
	tmpl, err := template.New("test").Parse(axx)
	f1, err := os.Create("x1.go")
	fmt.Println(err)
	defer f1.Close()
	err = tmpl.Execute(f1, x1)
	fmt.Println(tmpl, err)
}
