package main

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
	cli, err := CreateCli("http://localhost:8545")
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
	keystorePath  :=  "/Users/yekai/eth/data/keystore"
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
func 