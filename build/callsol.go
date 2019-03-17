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
func DeployPdbank() (common.Address, error) {
	auth, err := MakeAuth("0x70c53a4c94ccce9ce56effbfcb89b221f986cd41", "123")
	if err != nil {
		fmt.Println("failed to makeAuth", err)
		return common.HexToAddress(""), err
	}

	//common.Address, *types.Transaction, *Pdbank, error
	contractaddr, ts, _, err := contracts.DeployPdbank(auth,testclient,"pdj")
	if err != nil {
		fmt.Println("failed to deloy ",err)
		return common.HexToAddress(""), err
	}
	fmt.Println(ts.ChainId(), ts.Hash().Hex(), ts.Nonce())
	fmt.Println(contractaddr.Hex())
	return contractaddr, err
}

func CallDeposit(addr, pass string) (*types.Transaction, error) {

	instance, err := contracts.NewPdbank(common.HexToAddress(Config.Common.ContractAddr), testclient)
	if err != nil {
		fmt.Println("failed to get contract instance", err)
		return nil, err
	}
	auth, err := MakeAuth(addr, pass)
	if err != nil {
		fmt.Println("failed to makeAuth", err)
		return nil, err
	}
	auth.Value = big.NewInt(1000000000)
	ts,err := instance.Deposit(auth)
	if err != nil {
		fmt.Println("failed to call ", err)
		return nil, err
	}
	fmt.Println(ts.ChainId(), ts.Hash().Hex(), ts.Nonce())
	return ts , err
}

func CallWithdraw(addr, pass string) (*types.Transaction, error) {

	instance, err := contracts.NewPdbank(common.HexToAddress(Config.Common.ContractAddr), testclient)
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
	ts,err := instance.Withdraw(auth, big.NewInt(20000))
	if err != nil {
		fmt.Println("failed to call ", err)
		return nil, err
	}
	fmt.Println(ts.ChainId(), ts.Hash().Hex(), ts.Nonce())
	return ts , err
}

func CallBankName() (error) {
	instance, err := contracts.NewPdbank(common.HexToAddress(Config.Common.ContractAddr), testclient)
	if err != nil {
		fmt.Println("failed to get contract instance", err)
		return err
	}
	name , err := instance.BankName(nil)
	if err != nil {
		fmt.Println("failed to get Balances", err)
		return err
	}
	fmt.Println(name , err)
	return nil
}

func CallBalances() (error) {
	instance, err := contracts.NewPdbank(common.HexToAddress(Config.Common.ContractAddr), testclient)
	if err != nil {
		fmt.Println("failed to get contract instance", err)
		return err
	}
	data , err := instance.Balances(nil, common.HexToAddress("0x70c53a4c94ccce9ce56effbfcb89b221f986cd41"))
	if err != nil {
		fmt.Println("failed to get Balances", err)
		return err
	}
	fmt.Println(data , err)
	return nil
}

