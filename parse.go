package main

import (
	"io"
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
}

func ParseJson(filename string) {

}
