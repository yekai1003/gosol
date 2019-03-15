package main

import (
	"fmt"
	"io/ioutil"
	"log"
	_ "math/big"
	"os"

	"gosol/contracts"
	_ "gosol/toml"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/ethereum/go-ethereum/rpc"
)

var testclient *ethclient.Client

func init() {
	fmt.Println("init()  ∂∂connect eth begin")
	cli, err := CreateCli("http://localhost:8545")
	if err != nil {
		log.Panic("failed to connect to eth", err)
	}
	testclient = cli
	fmt.Println("init()  connect eth end")
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
	keystorePath := "/Users/yekai/eth/data/keystore"
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
func DeployPdbank() (common.Address, error) {
	var contractaddr common.Address
	auth, err := MakeAuth("0x70c53a4c94ccce9ce56effbfcb89b221f986cd41", "123")
	if err != nil {
		fmt.Println("failed to makeAuth", err)
		return contractaddr, err
	}
	auth.GasLimit = 300000
	fmt.Printf("%+v,,,%s\n", auth, auth.From.Hex())
	//common.Address, *types.Transaction, *Pdbank, error
	contractaddr, _, pb, err := contracts.DeployPdbank(auth, testclient)
	if err != nil {
		fmt.Println("failed to deloy ", err)
		return contractaddr, err
	}

	amount, err := pb.TotalAmount(nil)
	if err != nil {
		fmt.Println("failed to get amount", err)
	}
	fmt.Println(contractaddr.Hex(), amount)
	return contractaddr, err
}
func Run() {
	fmt.Println("deploy^^^^")
	if os.Args[1] == "deploy" {
		DeployPdbank()
	}
}
