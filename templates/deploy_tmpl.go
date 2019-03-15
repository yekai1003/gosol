package templates

const Main_tmpl = `package main

import (
	"fmt"
	"log"
	"math/big"
	"os"

	"gosol/contracts"
	"gosol/toml"
	"ioutil"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
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

const Deploy_sol_tmpl = `func {{.Func}}() (common.Address, error) {
	auth, err := MakeAuth("{{.FromAddr}}", "{{.Pass}}")
	if err != nil {
		fmt.Println("failed to makeAuth", err)
		return nil, err
	}

	//common.Address, *types.Transaction, *Pdbank, error
	contractaddr, _, _, err := contracts.{{.Params}}
	if err != nil {
		fmt.Println("failed to deloy ",err)
		return nil, err
	}
	return contractaddr, err
}
func Run() {
	fmt.Println("deploy^^^^")
	if os.Args[1] == "deploy" {
		{{.Func}}()
	}
}`

const Call_func_tmpl = `func CallAuthFunc(addr, pass, keystorePath string) (*types.Transaction, error) {

	instance, err := contracts.NewPdbank(common.HexToAddress(toml.Config.ContractAddr), testclient)
	if err != nil {
		fmt.Println("failed to get contract instance", err)
		return nil, err
	}
	auth, err := MakeAuth(addr, pass, keystorePath)
	if err != nil {
		fmt.Println("failed to makeAuth", err)
		return nil, err
	}
	auth.Value = 10000000000
	return instance.Deposit(auth)
}`

const Call_nogas_func_tmpl = `func CallNogasFunc(addr string) (*big.Int, error) {
	instance, err := contracts.NewPdbank(common.HexToAddress(toml.Config.ContractAddr), testclient)
	if err != nil {
		fmt.Println("failed to get contract instance", err)
		return nil, err
	}
	data, err := instance.Balances(common.HexToAddress(addr))
	if err != nil {
		fmt.Println("failed to get Balances", err)
		return nil, err
	}
	fmt.Println(data, err)
	return data, err
}`
