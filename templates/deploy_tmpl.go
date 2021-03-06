package templates

const Main_tmpl = `package main

import (
	"fmt"
	"log"
	"os"

	"gosol/contracts"
	"io/ioutil"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var testclient *ethclient.Client

func init() {
	cli, err := CreateCli("{{.Connstr}}")
	if err != nil {
		log.Panic("failed to connect to eth", err)
	}
	testclient = cli
}

func GetFileName(address, dirname string) (string, error) {

	data, err := ioutil.ReadDir(dirname)
	if err != nil {
		fmt.Println("read dir err", err)
		return "", err
	}
	for _, v := range data {
		if strings.Index(v.Name(), address) > 0 {
			//代表找到文件
			return v.Name(), nil
		}
	}

	return "", nil
}

//创建链接
func CreateCli(connstr string) (*ethclient.Client, error) {
	cli, err := ethclient.Dial(connstr)
	if err != nil {
		fmt.Println("failed to dial provide", err)
		return nil, err
	}
	return cli, err
}

//设置签名
func MakeAuth(addr, pass string) (*bind.TransactOpts, error) {
	keystorePath  :=  "{{.Keydir}}"
	fileName, err := GetFileName(string([]rune(addr)[2:]), keystorePath)
	if err != nil {
		fmt.Println("failed to GetFileName", err)
		return nil, err
	}

	file, err := os.Open(keystorePath + "/" + fileName)
	if err != nil {
		fmt.Println("failed to open file ", err)
		return nil, err
	}
	auth, err := bind.NewTransactor(file, pass)
	if err != nil {
		fmt.Println("failed to NewTransactor  ", err)
		return nil, err
	}
	return auth, err
}
`

const Deploy_sol_tmpl = `
func Deploy{{.ContractName}}() (common.Address, error) {
	auth, err := MakeAuth("{{.FromAddr}}", "{{.Pass}}")
	if err != nil {
		fmt.Println("failed to makeAuth", err)
		return common.HexToAddress(""), err
	}

	//common.Address, *types.Transaction, *Pdbank, error
	contractaddr, ts, _, err := contracts.{{.CallFunc}}
	if err != nil {
		fmt.Println("failed to deloy ",err)
		return common.HexToAddress(""), err
	}
	fmt.Println(ts.ChainId(), ts.Hash().Hex(), ts.Nonce())
	fmt.Println(contractaddr.Hex())
	return contractaddr, err
}

`

const Call_func_tmpl = `
func Call{{.FuncName}}(addr, pass string) (*types.Transaction, error) {

	instance, err := contracts.{{.Constructor}}(common.HexToAddress(Config.Common.ContractAddr), testclient)
	if err != nil {
		fmt.Println("failed to get contract instance", err)
		return nil, err
	}
	auth, err := MakeAuth(addr, pass)
	if err != nil {
		fmt.Println("failed to makeAuth", err)
		return nil, err
	}
	auth.Value = big.NewInt(0)
	ts,err := instance.{{.CallFunc}}
	if err != nil {
		fmt.Println("failed to call ", err)
		return nil, err
	}
	fmt.Println(ts.ChainId(), ts.Hash().Hex(), ts.Nonce())
	return ts , err
}

`

const Call_nogas_func_tmpl = `
func Call{{.FuncName}}() (error) {
	instance, err := contracts.{{.Constructor}}(common.HexToAddress(Config.Common.ContractAddr), testclient)
	if err != nil {
		fmt.Println("failed to get contract instance", err)
		return err
	}
	{{.RetFunc}} := instance.{{.CallFunc}}
	if err != nil {
		fmt.Println("failed to get Balances", err)
		return err
	}
	fmt.Println({{.RetFunc}})
	return nil
}

`

const Test_func_tmpl = `
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
`
const Test_funcinit_tmpl = `
var datas = []FuncsList{
%s
}


func main() {
	Run()
}

`

const Test_func_run_tmpl = `
func Run() {
	if len(os.Args) < 2 || os.Args[1] == "1" {
		Usage()
		os.Exit(0)
	}
	if os.Args[1] == "2" {
		%s()
	} `

const Test_func_run_tmpl2 = `else if os.Args[1] == "%d" {
		Call%s(%s, %s)
	} `

const Test_func_run_tmpl3 = `else if os.Args[1] == "%d" {
		Call%s()
	} `
