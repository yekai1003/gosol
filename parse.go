package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type NameType struct {
	ColName string `json:'name'`
	ColType string `json:'type'`
}

type ContractInfo struct {
	Constant        bool
	Inputs          []NameType
	Name            string
	Outputs         []NameType
	Payable         bool
	StateMutability string
	Type            string
}

func ParseJson(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Panic("faild to open file", filename, err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Panic("faild to read file", filename, err)
	}
	fmt.Println(string(data))
	var info []ContractInfo
	//info := make([]ContractInfo, 1)
	json.Unmarshal(data, &info)
	fmt.Printf("%+v\n", info[2])
}
func ParseDir(dirName string) ([]string, error) {
	var solFiles []string
	dirinfo, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Println("faild to read file", dirName, err)
		return nil, err
	}
	for _, oneInfo := range dirinfo {
		strrune := []rune(oneInfo.Name())
		if !oneInfo.IsDir() && len(strrune) > 3 {
			strfix := strrune[len(strrune)-3:]
			if string(strfix) == "sol" {
				solFiles = append(solFiles, oneInfo.Name())
			}
		}

		//fmt.Println(oneInfo.Name(), string(strfix), oneInfo.Size(), oneInfo)
	}
	return solFiles, nil
}

func ParseRun() {
	solfiles, err := ParseDir("sol")
	fmt.Println(solfiles, err)
	for _, solfile := range solfiles {
		fmt.Println(solfile)
		codeName, err := Compiler(solfile, "sol", "contracts")
		if err != nil {
			fmt.Println("failed to complie code", err)
			return
		}
		err = BuildAbi(codeName)
		if err != nil {
			fmt.Println("failed to build abi", err)
			return
		}
	}
}
